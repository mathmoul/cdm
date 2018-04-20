import React from 'react'
import isEmpty from 'lodash/isEmpty'
import PropType from 'prop-types'
import InlineError from '../messages/InlineError'
import { Form, Message, Button } from 'semantic-ui-react'

class ResetPasswordForm extends React.Component {
  state = {
    data: {
      token: this.props.token,
      password: '',
      passwordConfirmation: ''
    },
    loading: false,
    errors: {}
  }

  onChange = e => this.setState({
    data: {
      ...this.state.data,
      [e.target.name]: e.target.value
    }
  })

  onSubmit = e => {
    e.preventDefault()
    const errors = this.validate(this.state.data)
    this.setState({ errors })
    if (isEmpty(errors.password) && isEmpty(errors.passwordConfirmation)) {
      this.setState({ loading: true })
      this.props
        .submit(this.state.data)
        .catch(err =>
          this.setState({ errors: err.response.data.errors, loading: false })
        )
    }
  }

  validate = data => {
    const errors = {}
    if (!data.password) errors.password = 'Can\'t be blank'
    if (data.password !== data.passwordConfirmation) errors.password = 'Passwords must match'
    return errors
  }

  render() {
    const { loading, errors, data } = this.state
    return (

      <Form onSubmit={this.onSubmit} loading={loading}>
        {
          errors.global && <Message negative>
            <Message.Header>
              Something went wrong
            </Message.Header>
            <p>{errors.global}</p>
          </Message>
        }
        <Form.Field error={!!errors.password}>
          <label htmlFor="password">Password</label>
          <input
            type="password"
            id="password"
            name="password"
            placeholder="your new password"
            value={data.password}
            onChange={this.onChange}
          />
          {errors.password && <InlineError text={errors.password} />}
        </Form.Field>

        <Form.Field error={!!errors.passwordConfirmation}>
          <label htmlFor="password">Confirm your new password</label>
          <input
            type="password"
            id="passwordConfirmation"
            name="passwordConfirmation"
            placeholder=""
            value={data.passwordConfirmation}
            onChange={this.onChange}
          />
          {errors.passwordConfirmation && <InlineError text={errors.passwordConfirmation} />}
        </Form.Field>
        <Button primary>
          Reset
        </Button>
      </Form>
    )
  }
}

ResetPasswordForm.propTypes = {
  token: PropType.string.isRequired,
  submit: PropType.func.isRequired
}

export default ResetPasswordForm
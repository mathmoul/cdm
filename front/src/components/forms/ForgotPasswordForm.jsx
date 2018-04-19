import React from 'react'

import PropTypes from 'prop-types'

import Validator from 'validator'
import InlineError from '../messages/InlineError'
import isEmpty from 'lodash/isEmpty'


import { Button, Form, Message } from 'semantic-ui-react'

class ForgotPasswordForm extends React.Component {
  state = {
    data: {
      email: ''
    },
    errors: {},
    loading: false
  }

  onChange = (e) => this.setState({
    data: {
      ...this.state.data,
      [e.target.name]: e.target.value
    }
  })

  onSubmit = () => {
    const errors = this.validate(this.state.data)
    this.setState({errors})
    if (isEmpty(errors.email)) {
      this.setState({loading: true})
      this.props
        .submit(this.state.data)
        .catch(err =>
          this.setState({errors: err.response.data.errors, loading: false})
        )
    }
  }

  validate = (data) => {
    const errors = {
      email: '',
      password: ''
    }
    if (!Validator.isEmail(data.email)) {
      errors.email = 'Invalid email'
    }
    if (!data.password) {
      errors.password = 'Can\'t be blank'
    }
    return errors
  }

  render () {
    const {data, errors, loading} = this.state
    return (
      <Form onSubmit={this.onSubmit} loading={loading}>
        {!!errors.global && <Message>{errors.global}</Message>}
        <Form.Field error={!!errors.email}>
          <label htmlFor="email">Email</label>
          <input
            type="email"
            id="email"
            name="email"
            placeholder="example@example.com"
            value={data.email}
            onChange={this.onChange}
          />
          {errors.email && <InlineError text={errors.email}/>}
        </Form.Field>
        <Button primary>Reinitialisation du mot de passe</Button>
      </Form>
    )
  }
}

ForgotPasswordForm.propTypes = {
  submit: PropTypes.func.isRequired
}

export default ForgotPasswordForm
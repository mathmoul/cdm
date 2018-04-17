import React from 'react'
import PropTypes from 'prop-types'

import Validator from 'validator'
import InlineError from '../messages/InlineError'
import isEmpty from 'lodash/isEmpty'

import { Button, Form, Message } from 'semantic-ui-react'

class LoginForm extends React.Component {
  state = {
    data: {
      email: '',
      password: ''
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
    if (isEmpty(errors.email) && isEmpty(errors.password)) {
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
    console.log(errors)
    return (
      <Form onSubmit={this.onSubmit} loading={loading}>
        {
          errors.global && <Message negative>
            <Message.Header>
              Something went wrong
            </Message.Header>
            <p>invalid credentials</p>
          </Message>
        }
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

        <Form.Field error={!!errors.password}>
          <label htmlFor="password">Password</label>
          <input
            type="password"
            id="password"
            name="password"
            placeholder="example@example.com"
            value={data.password}
            onChange={this.onChange}
          />
          {errors.password && <InlineError text={errors.password}/>}
        </Form.Field>
        <Button primary>Login</Button>
      </Form>
    )
  }
}

LoginForm.propTypes = {
  submit: PropTypes.func.isRequired,
}

export default LoginForm

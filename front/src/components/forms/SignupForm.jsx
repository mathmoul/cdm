import React from 'react'
// import PropTypes from 'prop-types'
// import { connect } from 'react-redux'
import { Form, Button } from 'semantic-ui-react'
import isEmail from 'validator'
import InlineError from '../messages/InlineError'

class SignupForm extends React.Component {
  state = {
    data: {
      email: '',
      password: ''

    },
    loading: false,
    errors: {}
  }

  onSubmit = (e) => {
    e.preventDefault()
    const errors = this.validate(this.state.data)
    this.setState({errors})
    if (Object.keys(errors).length === 0) {
      this.setState({loading: true})
      this.props.submit(this.state.data)
        .catch(err => this.setState({errors: err.response.data.errors, loading: false}))
    }
  }

  validate = data => {
    const errors = {}
    if (!isEmail(data.email)) errors.email = 'Invalid Email'
    if (!data.password) errors.password = 'Can\'t be blank'
    return errors
  }

  onChange = (e) => {
    this.setState({
      ...this.state, data: {...this.state.data, [e.target.name]: e.target.value}
    })
  }

  render () {
    const {data, errors, loading} = this.state

    return (
      <Form onSubmit={this.onSubmit} loading={loading}>

        <Form.Field error={!!errors.email}>
          <label htmlFor="email">Email</label>
          <input
            type="email"
            id="email"
            name="email"
            placeholder="your email here"
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
        <Button primary>Signup</Button>
      </Form>
    )
  }
}

export default SignupForm
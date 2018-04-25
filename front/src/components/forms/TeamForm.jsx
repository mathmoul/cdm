import React from 'react'
import InlineError from '../messages/InlineError'
import {Button, Dropdown, Form, Message} from 'semantic-ui-react'

import coutryOptions from "../../Semantic_utils/FlagTab"

class TeamForm extends React.Component {
    state = {
        name: '',
        flag: '',
        value: "",
        searchQuery: "",
        loading: false,
        errors: {}
    }


    onChange = e => this.setState({
        data: {
            ...this.state.data,
            [e.target.name]: e.target.value
        }
    })

    onDropdownChange = (e, data) => {
        console.log(data);
        this.setState({flag: data.value})
    }

    onSubmit = e => {
        // e.preventDefault()
        // const errors = this.validate(this.state.data)
        // this.setState({errors})
        // if (isEmpty(errors.password) && isEmpty(errors.passwordConfirmation)) {
        //     this.setState({loading: true})
        //     this.props
        //         .submit(this.state.data)
        //         .catch(err =>
        //             this.setState({errors: err.response.data.errors, loading: false})
        //         )
        // }
    }

    // validate = data => {
    //     const errors = {}
    //     if (!data.password) errors.password = 'Can\'t be blank'
    //     if (data.password !== data.passwordConfirmation) errors.password = 'Passwords must match'
    //     return errors
    // }

    render() {
        const {loading, errors, name, flag} = this.state;
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
                <Form.Field error={!!errors.name}>
                    <label htmlFor="password">Nom</label>
                    <input
                        type="text"
                        id="name"
                        name="name"
                        placeholder="Nom du pays"
                        value={name}
                        onChange={this.onChange}
                    />
                    {errors.name && <InlineError text={errors.name}/>}
                </Form.Field>
                <Dropdown button icon="flag"
                          selection
                          options={coutryOptions}
                          value={flag}
                          onChange={this.onDropdownChange}
                />
                <Button primary>Save</Button>
            </Form>
        )
    }
}

TeamForm.propTypes = {};

export default TeamForm


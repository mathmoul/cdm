import React from 'react'
import PropTypes from 'prop-types'

import { Form, Button, Segment, Grid, Image } from 'semantic-ui-react'

import InlineError from "./../messages/InlineError"

class BookForm extends React.Component {
    state = {
        data: {
            id: this.props.book.id,
            title: this.props.book.title,
            author: this.props.book.author,
            cover: this.props.book.covers[0],
            pages: this.props.book.pages,
        },
        covers: this.props.book.covers,
        loading: false,
        errors: {}
    }

    onChange = e =>
        this.setState({
            ...this.state, data: {
                ...this.state.data, [e.target.name]: e.target.value
            }
        })


    onChangeNumber = e =>
        this.setState({
            ...this.state, data: {
                ...this.state.data, [e.target.name]: parseInt(e.target.value, 10)
            }
        })

    onSubmit = e => {
        e.preventDefault()
        const errors = this.validate(this.state.data)
        this.setState({ errors })
        if (Object.keys(errors).length === 0) {
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
        if (!data.title) errors.title = "can't be blank"
        return errors
    }


    componentWillReceiveProps(props) {
        
        this.setState({
            data: {
                id: props.book.id,
                title: props.book.title,
                author: props.book.author,
                cover: props.book.covers[0],
                pages: props.book.pages,
            },
            covers: props.book.covers,
        })
    }

    render() {
        const { errors, data } = this.state
        return (
            <Segment>
                <Form onSubmit={this.onSubmit} loading={this.loading}>
                    <Grid column={2}>
                        <Grid.Row>
                            <Grid.Column>
                                <Form.Field error={!!errors.title}>
                                    <label htmlFor="title">Book title</label>
                                    <input
                                        type="text"
                                        id="title"
                                        name="title"
                                        placeholder="Title"
                                        value={data.title}
                                        onChange={this.onChange}
                                    />
                                    {errors.title && <InlineError text={errors.title} />}
                                </Form.Field>


                                <Form.Field error={!!errors.author}>
                                    <label htmlFor="author">Book author</label>
                                    <input
                                        type="text"
                                        id="author"
                                        name="author"
                                        placeholder="Author"
                                        value={data.author}
                                        onChange={this.onChange}
                                    />
                                    {errors.author && <InlineError text={errors.author} />}
                                </Form.Field>

                                <Form.Field error={!!errors.pages}>
                                    <label htmlFor="title">Pages</label>
                                    <input
                                        type="number"
                                        id="pages"
                                        name="pages"
                                        value={data.pages}
                                        onChange={this.onChangeNumber}
                                    />
                                    {errors.pages && <InlineError text={errors.pages} />}
                                </Form.Field>

                            </Grid.Column>
                        </Grid.Row>
                        <Grid.Row>
                            <Grid.Column>
                                <Image size="small" src={data.cover} />
                            </Grid.Column>
                        </Grid.Row>
                    </Grid>
                    <Button primary>Save</Button>
                </Form>
            </Segment>
        )
    }
}

BookForm.propTypes = {
    submit: PropTypes.func.isRequired,
    book: PropTypes.shape({
        id: PropTypes.string.isRequired,
        title: PropTypes.string.isRequired,
        author: PropTypes.string.isRequired,
        covers: PropTypes.arrayOf(PropTypes.string.isRequired).isRequired,
        pages: PropTypes.number.isRequired,
    }).isRequired
}

export default BookForm
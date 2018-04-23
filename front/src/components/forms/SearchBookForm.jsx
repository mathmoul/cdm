import React, { Component } from "react";
import PropTypes from 'prop-types'
import axios from "axios";

import { Form, Dropdown } from "semantic-ui-react";

class SearchBookForm extends Component {
  state = {
    searchQuery: "",
    value: null,
    loading: false,
    options: [],
    books: {}
  };

  onChange = (e, data) => {
    console.log("dataaaaa  ========>", data)
    this.setState({ searchQuery: this.state.books[data.value].title })
    this.props.onBookSelect(this.state.books[data.value])
  }

  onSearchChange = (e, data) => {
    console.log(this.state);
    console.log(data);
    clearTimeout(this.timer);
    this.setState({ searchQuery: data.searchQuery });
    this.timer = setTimeout(this.fetchOptions, 1000);
  };

  fetchOptions = () => {
    console.log(this.state.searchQuery);
    if (!this.state.searchQuery) return;
    this.setState({ loading: true });
    axios
      .get(`/api/books/search?q=${this.state.searchQuery}`)
      .then(res => res.data.books)
      .then(books => {
        const options = []
        const booksHash = {}
        console.log(books)
        books.forEach(book => {
          booksHash[book.id] = book
          options.push({
            key: book.id,
            value: book.id,
            text: book.title
          })
        });
        console.log(options, booksHash)
        this.setState({ loading: false, options, books: booksHash })
      })
  };

  // TODO update semantic ui

  render() {
    const { searchQuery, value } = this.state;
    console.log(this.state)
    return (
      <Form>
        <Dropdown
          search
          fluid
          placeholder="Search for a book by title"
          value={value}
          searchQuery={searchQuery}
          onSearchChange={this.onSearchChange}
          options={this.state.options}
          loading={this.state.loading}
          onChange={this.onChange}
        />
      </Form>
    );
  }
}

SearchBookForm.propTypes = {
  onBookSelect: PropTypes.func.isRequired
}

export default SearchBookForm;

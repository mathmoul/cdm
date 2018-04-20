import React, { Component } from "react";

import axios from "axios";

import { Form, Dropdown } from "semantic-ui-react";

class SearchBookForm extends Component {
  state = {
    searchQuery: "",
    value: null,
    loading: false,
    options: [
      {
        key: 1,
        value: 1,
        text: "test"
      },
      {
        key: 2,
        value: 2,
        text: "test2"
      }
    ],
    books: {}
  };

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
      .then(res => res.data.books);
  };

  // TODO update semantic ui

  render () {
    const { searchQuery, value } = this.state;
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
        />
      </Form>
    );
  }
}

export default SearchBookForm;

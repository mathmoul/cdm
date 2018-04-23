import React, { Component } from "react";
import { connect } from "react-redux";

import { Segment } from "semantic-ui-react";
import SearchBookForm from "../forms/SearchBookForm";

import BookForm from '../forms/BookForm'

class NewBookPage extends Component {
  state = {
    book: null
  };

  onBookSelect = book => {
    this.setState({ book })
  }

  addBook = book => console.log(book)

  render() {
    return (
      <Segment>
        <h1>Add New Book to your collection</h1>
        <SearchBookForm onBookSelect={this.onBookSelect} />

        {this.state.book &&
          (<BookForm submit={this.addBook} book={this.state.book} />
          )}
      </Segment>
    );
  }
}

function mapStateToProps(state) {
  return {};
}

export default connect(mapStateToProps)(NewBookPage);

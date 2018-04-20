import React, { Component } from "react";
import { connect } from "react-redux";

import { Segment } from "semantic-ui-react";
import SearchBookForm from "../forms/SearchBookForm";

class NewBookPage extends Component {
  state = {
    book: null
  };

  render() {
    return (
      <Segment>
        <h1>Add New Book to your collection</h1>
        <SearchBookForm />
      </Segment>
    );
  }
}

function mapStateToProps(state) {
  return {};
}

export default connect(mapStateToProps)(NewBookPage);

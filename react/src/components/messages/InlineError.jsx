import PropTypes from "prop-types";
import * as React from "react";

const InlineError = ({text}) => <span style={{color: "#ae5856"}}>{text}</span>;

InlineError.propTypes = {
    text: PropTypes.string.isRequired,
};

export default InlineError;

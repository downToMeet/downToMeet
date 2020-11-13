import React from "react";
import PropTypes from "prop-types";

function Meetup({ id }) {
  return <>Information about meetup with id {id}</>;
}

Meetup.propTypes = {
  id: PropTypes.string.isRequired,
};

export default Meetup;

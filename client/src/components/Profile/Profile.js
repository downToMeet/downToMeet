import React from "react";
import PropTypes from "prop-types";

function Profile({ id }) {
  return <>Profile of user with id {id}</>;
}

Profile.propTypes = {
  id: PropTypes.string.isRequired,
};

export default Profile;

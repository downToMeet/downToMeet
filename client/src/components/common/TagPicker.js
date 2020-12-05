import React from "react";
import { TextField } from "@material-ui/core";
import AutoComplete from "@material-ui/lab/Autocomplete";
import PropTypes from "prop-types";
import { tagOptions } from "../../constants";

export default function TagPicker({ tags, setTags }) {
  return (
    <AutoComplete
      multiple
      value={tags}
      onChange={(event, newValue) => setTags(newValue)}
      variant="outlined"
      options={tagOptions}
      renderInput={(params) => (
        <TextField
          // eslint-disable-next-line react/jsx-props-no-spreading
          {...params}
          variant="outlined"
          label="Look for tags"
        />
      )}
      style={{ flex: 1 }}
    />
  );
}

TagPicker.propTypes = {
  tags: PropTypes.arrayOf(PropTypes.string.isRequired).isRequired,
  setTags: PropTypes.func.isRequired,
};

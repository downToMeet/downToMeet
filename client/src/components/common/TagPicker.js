import React from "react";
import { TextField } from "@material-ui/core";
import AutoComplete from "@material-ui/lab/Autocomplete";
import PropTypes from "prop-types";
import { tagOptions } from "../../constants";

/**
 * Dropdown menu for picking interest tags related to a meetup.
 *
 * Used to pick meetup tags in [CreateMeetup](#createmeetup) and [Search](#search).
 */
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
          required
          variant="outlined"
          label="Look for tags"
        />
      )}
      style={{ flex: 1 }}
    />
  );
}

TagPicker.propTypes = {
  /** List of tags picked. */
  tags: PropTypes.arrayOf(PropTypes.string.isRequired).isRequired,
  /** Setter for `tags`, generated from `useEffect()`. */
  setTags: PropTypes.func.isRequired,
};

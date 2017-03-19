import React, { Component } from 'react';
import { Pill } from 'elemental';

class HashTags extends Component {
  render() {
    const { hashtags, handleClearHashtag, childClassName } = this.props;
    return (
      <div>
        {
          hashtags.map((hashtag, i) => (
          <Pill
            key={i}
            label={hashtag}
            type={childClassName}
            onClear={hashtag => handleClearHashtag(hashtag)}
          />
        ))
      }
      </div>
    );
  }
}

HashTags.propTypes = {
  childClassName: React.PropTypes.string.isRequired,
  hashtags: React.PropTypes.array.isRequired,
  handleClearHashtag: React.PropTypes.func.isRequired
};

HashTags.defaultProps = {
  hashtags: []
};

export default HashTags;

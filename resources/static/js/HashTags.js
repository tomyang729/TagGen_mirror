import React, { Component } from 'react';
import { Pill } from 'elemental';

import '../css/HashTags.css';

class HashTags extends Component {
  render() {
    const { hashtags, handleClearHashtag, childClassName } = this.props;
    return (
        <div className="HashTags">
        {
          hashtags.map((hashtag, i) => (
          <Pill className="pill"
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
  childClassName: React.PropTypes.string,
  hashtags: React.PropTypes.array,
  handleClearHashtag: React.PropTypes.func
};

HashTags.defaultProps = {
  hashtags: ["dog", "cat", "house", "computer", "bank", "banana"]
};

export default HashTags;

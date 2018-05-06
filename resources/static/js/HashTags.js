import React, { Component } from 'react';
import { Pill } from 'elemental';

import '../css/HashTags.css';

class HashTags extends Component {
  _handleClearHashTag(e, hashtag) {
    e.preventDefault();
    this.props.handleClearHashtag(hashtag);
  }

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
            onClear={e => this._handleClearHashTag(e, hashtag)}
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

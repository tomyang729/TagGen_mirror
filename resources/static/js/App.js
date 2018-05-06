import React, { Component } from 'react';
import logo from '../images/hashtag1.png';
import { Button } from 'elemental';
import HashTags from './HashTags';
import FileUpload from './FileUpload';
import axios from 'axios';

import '../css/App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      hashtags: []
    };
    this._handleClearHashtag = this._handleClearHashtag.bind(this);
    this._getHashTags = this._getHashTags.bind(this);
  }

  _handleClearHashtag(hashtag) {
    this.setState({
      hashtags: this.state.hashtags.filter(ht => ht !== hashtag)
    })
  }

  _getHashTags(imageURI) {
    let url = 'http://localhost:5050/getTags/';
    
    axios({
      method: 'post',
      url: url,
      data: {image: imageURI}
    })
    .then(response => {
      this.setState({
        hashtags: response.data
      });
    })
    .catch(error => {
      alert("Sorry something wrong. Please try again!\n" + error);
    });
  }

  render() {
    return (
      <div className="App">
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/latest/css/bootstrap.min.css"></link>
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1><b>Welcome to #TagGen!</b></h1>
        </div>
        <div className="container-fluid main-container">
          <FileUpload getHashTags={this._getHashTags} />
          <HashTags
            childClassName="primary"
            hashtags={this.state.hashtags}
            handleClearHashtag={this._handleClearHashtag}
          />
        </div>
      </div>
    );
  }
}

export default App;

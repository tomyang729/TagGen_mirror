import React, { Component } from 'react';
import logo from '../images/logo.svg';
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

  _getHashTags(imageURI, file) {
    let url = 'http://localhost:5050/getTags/';
    let formdata = new FormData();
    let image = {
      uri: imageURI,
      name: file.name,
      type: file.type || 'image-type/png'
    };
    formdata.append('image', image);
    axios({
      method: 'post',
      url: url,
      data: formdata
    })
    .then(response => {
      console.log(response);
      this.setState({
        hashtags: response
      });
    })
    .catch(error => {
      alert("Sorry something wrong. Please try again!\n" + error);
    });
  }

  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h2>Welcome to React</h2>
        </div>
        <FileUpload getHashTags={this._getHashTags}/>
        <HashTags
          childClassName="primary"
          hashtags={this.state.hashtags}
          handleClearHashtag={this._handleClearHashtag}
        />
      </div>
    );
  }
}

export default App;

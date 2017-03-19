import React, { Component } from 'react';
import { Button, Spinner } from 'elemental';
import placeholder from '../images/placeholder.jpg';

import '../css/FileUpload.css';

class FileUpload extends Component {
  constructor(props) {
    super(props);
    this.state = {
      file: {},
      dataURI: null,
      loading: false
    };
    this.triggerFileBrowser = this.triggerFileBrowser.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.cancelUpload = this.cancelUpload.bind(this);
  }

  componentDidMount () {
		this.refs.fileInput.addEventListener('click', function () {
			this.value = '';
		}, false);
	}

	triggerFileBrowser () {
		this.refs.fileInput.click();
	}

	handleChange (e) {
		var reader = new FileReader();
		var file = e.target.files[0];

		reader.readAsDataURL(file);

		reader.onloadstart = () => {
			this.setState({
				loading: true,
			});
		};
		reader.onloadend = (upload) => {
			this.setState({
				loading: false,
				file: file,
				dataURI: upload.target.result
			});
		};
	}

	cancelUpload (e) {
		this.setState({
			dataURI: null,
			file: {},
		});
	}

  render() {
    let { dataURI, file } = this.state;
    let $imagePreview = null;
    if (dataURI) {
      $imagePreview = (<img src={dataURI} className="picture"/>);
    } else {
      $imagePreview = (<div className="previewText"><h3>Simply upload a picture and get top most popular HashTags related
			to your picture!</h3><img src={placeholder}></img></div>);
    }

		let buttons = dataURI ? (
			<div className="previewComponent">
				<Button onClick={this.triggerFileBrowser} type="link-cancel" disabled={this.state.loading}>
					{this.state.loading && <Spinner />}
					{'Change File'}
				</Button>
				<Button onClick={this.cancelUpload} type="link-cancel" disabled={this.state.loading}>Cancel</Button>
			</div>
		) : (
				<Button type="button" className="btn btn-info button-general" onClick={this.triggerFileBrowser} disabled={this.props.disabled || this.state.loading}>
					{this.state.loading ? <Spinner /> : null}
				{'Upload File'}
			</Button>
		);

		return (
			<div>
        <div className="imgPreview">
          {$imagePreview}
        </div>
				{buttons}
				<input style={{ display: 'none' }} type="file" ref="fileInput" onChange={this.handleChange} />
        <Button type="button" className="btn btn-lg btn-warning button-general" onClick={(dataURI, file) => this.props.getHashTags(dataURI, file)}>Get #Hashtags!</Button>
			</div>
		);
  }
}

FileUpload.propTypes = {
  disabled: React.PropTypes.bool,
  getHashTags: React.PropTypes.func.isRequired
}

export default FileUpload;

import React, { Component } from 'react';
import { Button, Spinner } from 'elemental';

import '../css/FileUpload.css';

class FileUpload extends Component {
  constructor(props) {
    super(props);
    this.state = {
      file: {},
      dataURI: null,
      loading: false
    };
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
			if (typeof this.props.onChange === 'function') {
				this.props.onChange(e, {
					file: file,
					dataURI: upload.target.result
				});
			}
		};
	}

	cancelUpload (e) {
		this.setState({
			dataURI: null,
			file: {},
		});
		this.props.onChange(e, null);
	}

  render() {
    let { dataURI, file } = this.state;

    let { dataURI } = this.state;
    let $imagePreview = null;
    if (dataURI) {
      $imagePreview = (<img src={dataURI} />);
    } else {
      $imagePreview = (<div className="previewText">Please select an Image for Preview</div>);
    }

		let buttons = dataURI ? (
			<div className="previewComponent">
				<div className="fileUploadContent">
					<div className="fileUploadButtons">
						<Button onClick={this.triggerFileBrowser} disabled={this.state.loading}>
							{this.state.loading && <Spinner />}
							{'Change File'}
						</Button>
						<Button onClick={this.cancelUpload} type="link-cancel" disabled={this.state.loading}>Cancel</Button>
					</div>
				</div>
			</div>
		) : (
			<Button onClick={this.triggerFileBrowser} disabled={this.props.disabled || this.state.loading}>
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
				<input style={{ display: 'none' }} type="file" ref="fileInput" onChange={this.handleChange} {...props} />
        <Button onClick={this.props.getHashTags(dataURI, file)}>Generate Hashtags!</Button>
			</div>
		);
  }
}

FileUpload.propTypes = {
  disabled: React.PropTypes.bool,
	onChange: React.PropTypes.func,
  getHashTags: React.PropTypes.func.isRequired
}

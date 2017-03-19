const path = require('path');

//const CleanCSSPlugin = require("less-plugin-clean-css");

module.exports = {
  entry: './static/js/index.js',
  output: {
    path: path.resolve('static/js/bin'),
    filename: 'bundle.js'
  },
  module: {
    loaders: [
      { test: /\.(css|less)$/, loaders: ["style-loader", "css-loader", "less-loader"] },
      {
        exclude: [
          /\.html$/,
          /\.(js|jsx)$/,
          /\.css$/,
          /\.json$/,
          /\.bmp$/,
          /\.gif$/,
          /\.jpe?g$/,
          /\.png$/,
        ],
        loader: 'file-loader',
        options: {
          name: 'static/images/[name].[ext]',
        },
      },
      {
        test: [/\.bmp$/, /\.gif$/, /\.jpe?g$/, /\.png$/],
        loader: 'url-loader',
        options: {
          limit: 10000,
          name: 'static/images/[name].[ext]',
        },
      },
      { test: /\.(js|jsx)$/, loader: 'babel-loader', exclude: /node_modules/ },
    ]
  }
}

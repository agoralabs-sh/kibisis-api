const HtmlWebpackPlugin = require('html-webpack-plugin');
const HtmlInlineScriptPlugin = require('html-inline-script-webpack-plugin');
const path = require('path');
const TSConfigPathsPlugin = require('tsconfig-paths-webpack-plugin');
const webpack = require('webpack');

const distPath = path.resolve(__dirname, 'dist');
const mode = 'production';
const srcPath = path.resolve(__dirname, 'src');

module.exports = [
  /**
   * index - builds the html and inlines the scripts
   */
  {
    entry: {
      index: path.resolve(srcPath, 'index.tsx'),
    },

    mode,

    module: {
      rules: [
        {
          test: /\.css$/,
          use: [
            { loader: 'style-loader' },
            { loader: 'css-loader' },
          ]
        },
        {
          loader: 'handlebars-loader',
          test: /\.hbs$/,
        },
        {
          exclude: /node_modules/,
          test: /\.tsx?$/,
          loader: 'ts-loader',
          options: {
            compilerOptions: {
              jsx: 'react',
              lib: ['DOM', 'ESNext'],
              target: 'ES5'
            },
            transpileOnly: true,
          },
        },
      ]
    },

    name: 'index',

    output: {
      filename: '[name].js',
      path: distPath,
    },

    plugins: [
      new webpack.ProvidePlugin({
        Buffer: ['buffer', 'Buffer'],
        process: 'process/browser',
      }),
      new HtmlWebpackPlugin({
        chunks: ['index'],
        filename: 'index.html',
        template: path.resolve(srcPath, 'index.hbs'),
        title: 'Swagger UI',
      }),
      new HtmlInlineScriptPlugin(),
    ],

    resolve: {
      extensions: ['.js', '.ts', '.tsx'],
      fallback: {
        buffer: require.resolve('buffer'),
        stream: require.resolve('stream-browserify'),
      },
      plugins: [
        new TSConfigPathsPlugin(),
      ],
    },
  },

  /**
   * main - the function handler that serves up the html.
   */
  {
    entry: {
      main: path.resolve(srcPath, 'main.ts'),
    },

    mode,

    module: {
      rules: [{
        exclude: /node_modules/,
        test: /\.ts?$/,
        loader: 'ts-loader',
        options: {
          compilerOptions: {
            lib: ['ESNext'],
            target: 'ES6',
          },
          transpileOnly: true,
        },
      }]
    },

    name: 'main',

    output: {
      filename: '[name].js',
      path: distPath,
    },

    resolve: {
      extensions: ['.js', '.ts'],
      plugins: [
        new TSConfigPathsPlugin(),
      ],
    },

    target: 'node18',
  },
];

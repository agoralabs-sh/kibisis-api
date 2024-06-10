import 'dotenv/config';
import HtmlWebpackPlugin from 'html-webpack-plugin';
import HtmlInlineScriptPlugin from 'html-inline-script-webpack-plugin';
import { resolve } from 'node:path';
import TSConfigPathsPlugin from 'tsconfig-paths-webpack-plugin';
import { Configuration, DefinePlugin, ProvidePlugin } from 'webpack';

const distPath = resolve(__dirname, 'dist');
const srcPath = resolve(__dirname, 'src');

const config: Configuration = {
  entry: {
    index: resolve(srcPath, 'index.tsx'),
  },

  mode: 'production',

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
          configFile: resolve(__dirname, 'tsconfig.json'),
          logLevel: 'info',
          transpileOnly: true,
        },
      },
    ]
  },

  output: {
    filename: '[name].js',
    path: distPath,
  },

  plugins: [
    new DefinePlugin({
      __SWAGGER_SPEC_URL__: JSON.stringify(process.env.SWAGGER_SPEC_URL),
    }),
    new ProvidePlugin({
      Buffer: ['buffer', 'Buffer'],
      process: 'process/browser',
    }),
    new HtmlWebpackPlugin({
      chunks: ['index'],
      filename: 'index.html',
      template: resolve(srcPath, 'index.hbs'),
      title: 'Kibisis API',
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
};

export default config;

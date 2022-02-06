const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');

const BUILD_DIR = path.resolve(__dirname, 'build');

module.exports = {
    mode: 'none',
    entry: {
        app: path.join(__dirname, 'src', 'index.tsx')
    },
    output: {
        path: BUILD_DIR,
        filename: '[name].js',
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
                exclude: '/node_modules/'
            },
            {
                test: /\.(s*)css$/,
                use: ["style-loader", "css-loader", "sass-loader"],
            },
            {
                test: /\.(png|svg|jpg|gif)$/,
                loader: 'file-loader',
                exclude: '/node_modules/',
                options: { outputPath: '/img' }
            },
            {
                test: /\.(woff|woff2|eot|ttf|otf)$/i,
                loader: 'file-loader',
                exclude: '/node_modules/',
                options: { outputPath: '/fonts' }
            },
        ],
    },
    resolve: {
        alias: {
            "@components": path.resolve(__dirname, 'src/components/'),
            "@src": path.resolve(__dirname, 'src/'),
            "@features": path.resolve(__dirname, 'src/features/'),
            "@models": path.resolve(__dirname, 'src/models/'),
            "@helpers": path.resolve(__dirname, 'src/helpers/'),
            "@consts": path.resolve(__dirname, 'src/consts/'),
            "@views": path.resolve(__dirname, 'src/views/'),
            "@theme": path.resolve(__dirname, 'src/theme/'),
        },
        extensions: ['.ts', '.tsx', '.js']
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: path.join(__dirname, 'src', 'index.html')
        }),
        new CopyWebpackPlugin({
            patterns: [
                { from: './public/img', to: './img'},
                { from: './public/icons', to: './icons'},
                { from: './public/fonts', to: './fonts'}
            ]
        })
    ],
    devServer: {
        noInfo: true,
        hot: true,
        port: 3001,
        historyApiFallback: true,
        contentBase: [path.join(__dirname, 'build')],
    },
}


let mix = require('laravel-mix');

/*
 |--------------------------------------------------------------------------
 | Mix Asset Management
 |--------------------------------------------------------------------------
 |
 | Mix provides a clean, fluent API for defining some Webpack build steps
 | for your Laravel application. By default, we are compiling the Sass
 | file for your application, as well as bundling up your JS files.
 |
 */

mix.js("resources/assets/js/app.js", "public/js")
    .sass("resources/assets/sass/app.scss", "public/css")
    .sass("resources/assets/sass/default.scss", "public/css")
    .sass("resources/assets/sass/custom.scss", "public/css")
    .copy("resources/assets/img/!(*.afdesign)", "public/img")
    .sourceMaps();


const webpack = require("webpack");

module.exports = {
 configureWebpack: {
   plugins: [
     new webpack.ProvidePlugin({
       $: "jquery",
       jQuery: "jquery"
     })
   ]
 },
 css: {
  loaderOptions: {
    less: {
      lessOptions: {
        modifyVars: {
          'primary-color': '#1DA57A',
          'link-color': '#1DA57A',
          'border-radius-base': '12px',
        },
        javascriptEnabled: true,
        typescriptEnabled:true,
      },
    },
  },
},
};
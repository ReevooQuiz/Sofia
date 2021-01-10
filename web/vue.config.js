module.exports = {
    css: {
      loaderOptions: {
        less: {
          lessOptions: {
            modifyVars: {
              'primary-color': '#88d5d1',
              'link-color': '#88d5d1',
              'border-radius-base': '5px',
            },
            javascriptEnabled: true,
          },
        },
        
      },
    },
    lintOnSave: false,   // 关闭代码验证
    devServer: {
      port:9099,
    }
  };

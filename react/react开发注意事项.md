# react 学习总结

## 开发当中的注意事项：
  1. 一旦安装了 React，强烈建议您将它配置为生产构建进程 ，以确保您在生产环境中，使用最新版本的 React。
  2. 特别提示：对于国内用户，我们强烈建议您设置npm仓库的地址为淘宝镜像。设置方法是：
    - npm config set registry https://registry.npm.taobao.org/
  3. 每个文件当中都要引入React和Component组件，因为 JSX 编译后会调用 React.createElement 方法，所以在你的 JSX 代码中必须首先声明 React 变量。
    - import React, {Component}from 'react'
  4. 书写 JSX 的时候一般都会带上换行和缩进，这样可以增强代码的可读性。与此同时，我们同样推荐在 JSX 代码的外面扩上一个小括号，这样可以防止 分号自动插入的bug.
  5. 因为 JSX 的特性更接近 JavaScript 而不是 HTML , 所以 React DOM 使用 camelCase 小驼峰命名 来定义属性的名称，而不是使用 HTML 的属性名称。例如：
    - class 变成了 className，而 tabindex 则对应着 tabIndex.
  6. React 当中的元素事实上是普通的对象,元素事实上只是构成组件的一个部分.
  7. 在实际生产开发中, 大多数React应用只会调用一次 ReactDOM.render()
  8. 组件名称必须以大写字母开头。因为React 会将小写开头的标签名认为是 HTML 原生标签。且组件名称不支持表达式，但支持大写开头的变量。例如：
    - <div />表示一个DOM标签，但<Welcome />表示一个组件并限定了它的可用范围。
  9. 组件的返回值只能有一个根元素。
  10. Props的只读性。React是非常灵活的，但它也有一个严格的规则：
    <text style="color:red">所有的React组件必须像纯函数那样使用它们的props。</text>
  11. 记住：React 中的数据流是单向的，并在组件层次结构中向下传递。
  12. 请记住，代码是用来读的，这比写更重要
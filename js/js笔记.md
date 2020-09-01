# 重要知识点：

- 函数名.length     返回此函数形参的数量
- arguments.length  返回此函数实参的数量
- 函数柯里化（通用版）

    ```js
        // 通用版
        const curring = (fn, arr = []) => {
            // 函数的length 为函数形参的个数
            let len = fn.length
            return (...args) => {
                arr = arr.concat(args)
                return arr.length < len ? curring(fn, arr) : fn(...arr)
            }
        }
        // example
        const add = (a, b, c, d, e) => {
            return a + b + c + d + e
        }
        let r = curring(add)(1)(2)(3)(4)(5)
        console.log(r) // 15
    ```
package main

import(
    "fmt"
    "context"
)

func main() {
    ctx := context.Background()
    fmt.Println(ctx);
    fmt.Printf("The Context type is %T\n", ctx);
    fmt.Printf("The Context error is %T\n", ctx.Err())

    ctx, cancel := context.WithCancel(ctx)
    fmt.Println(ctx);
    fmt.Printf("The value of cancel variable %T\n", cancel);
    fmt.Printf("The type of cancel variable \n", cancel);
    fmt.Printf("The Context type is %T\n", ctx);
    fmt.Printf("The Context error is %T\n", ctx.Err())

    // do the cancelling
    cancel()

    fmt.Printf("After we cancel The value of cancel variable %T\n", cancel);
    fmt.Printf("After we cancel The type of cancel variable \n", cancel);
    fmt.Printf("After we cancel The Context type is %T\n", ctx);
    fmt.Printf("After we cancel The Context error is %T\n", ctx.Err())
}

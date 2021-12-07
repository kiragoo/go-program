# 链表笔记

## 访问模式

* 迭代

```golang

type ListNode struct{
    Val int
    Next *ListNode
}

func traverse(head *ListNode) {
    for cur:=head;cur != nil;cur = cur.Next{
        # 处理逻辑
    }
}
```

* 递归

```golang
func traverse(head *ListNode) {
    if head == nil {
        #  处理逻辑
    }
    traverse(head.Next)
    # 具体的处理逻辑
}
# Prompt 


## Prompt 是什么
随着人工智能领域的迅速发展，Prompt Engineering 已成为一门备受关注的新兴技术。
Prompt 是指用于引导机器学习模型生成符合预期输出的文本或代码片段 。
编写高效的 Prompt 对于提高模型效果至关重要。本文将探讨如何高效编写 Prompt。

一个高效的 Prompt 应由 Instruction 和 Question 两部分组成。
在使用 openAPI 调用时，通常将 Instruction 通过 system 传递，将 Question 通过 user 传递。

## 背景信息
本文档旨在为视频文本分类任务提供一个更有效的 prompt。
原始的 prompt 已经定义了用户需要从一段关于探险的视频文本中提取出3到4个观众可能感兴趣的类别。
我们将通过一些技巧来改进这个 prompt，以获得更准确的结果。

## Instruction
Instruction 通常由 **context** 和 **steps** 组成，这两者并非缺一不可的，也不存在清晰的界限。 
- Context 的格式是：
You are an agent/assistant of xxx. To xxx, you should follow next steps:
你是一个用来xxx的xxx，为了达到xxx目的，你需要遵循以下步骤：

- 关于 steps，需要根据具体问题来编写，通常我们会使用 Markdown 的列表形式：
  - step1
  - step2
  - step3

### 对于笔者来说一个好的prompt必须包含：
- 上下文（Context，背景）
- 目标（Objective，要做什么）
- 响应 （Response，输出）
- 示例（Example，例子）
> 在笔者认为 **Example** 的效果是最显著的。因为 **Example** 可以帮助模型理解问题的要求和目标，从而生成更符合要求的答案。


## 实战
### 第一版 Prompt

Input:
```plaintext
Help me extract the video keywords
```
Output:
```
**Keywords:哀牢山、独自进出、山区公司、寻找一级矿、白腹锦鸡、熊、裂缝河流、祖母绿矿脉、皮艇、大鱼、柔蓝色矿石**

1. **哀牢山**：整个故事发生的地点，是一个充满危险和神秘的地方。
2. **独自进出**：强调博主与国家地质科考队对比下独自进入哀牢山两次的特殊行为。
3. **山区公司**：是博主进入哀牢山寻找一级矿的原因，与故事的起因相关。
4. **寻找一级矿**：明确了博主进入哀牢山的目的。
5. **白腹锦鸡**：进山捡到的巨大鸟毛可能所属的动物，增加了山中生物的元素。
6. **熊**：深山野林中遇到的危险动物，是对山中危险的一种体现。
7. **裂缝河流**：在探险过程中的一个阻碍，深度探测不到底，影响前行路线。
8. **祖母绿矿脉**：博主对河堤可能存在物质的推测，与寻找矿有关。
9. **皮艇**：第二次进入时携带的工具，在探险过程中起到过河等作用。
10. **大鱼**：在河中遇到的生物，增加探险中的未知与惊险元素。
11. **柔蓝色矿石**：最终到达目的地发现的矿石，是博主勇闯哀牢山的目标成果。
```

> 显然并不符合我们的要求，我们的任务是只需要提取视频关键字，来为后期给视频分类。


### 第二版本 Prompt

Input:
```plaintext
Extract 3 to 4 keywords from this video text about to help with video classification.
```
Output:
```
**Keywords:哀牢山、寻找矿石、独自进出、河中不明物**
```

> 从这里似乎比前版本更符合我们的要求，但是在后期测试时会混合着中文和英文，导致分类结果不准确。

> 我们可以给稍加微的输出示例：
> 美食 | 旅行 | 阅读

### 第三版 Prompt
Input:
```plaintext
You will be provided with a block of text which is the content of a video, and your task is to give 5 keyword in Simplified Chinese to the video to attract audience.
For example, 美食 | 旅行 | 阅读
```
Output:
```paintext
哀牢山 | 独自探险 | 神秘矿脉 | 河中不明物 | 柔蓝色矿石
```

> 到这里可以看到，我们已经基本符合我们的要求。针对于这样的输出我们可以通过固定输出分割符，来转换成我们的需要的输出格式。但是在对于生产环境下，我们还需要进一步的校验操作，因为ai是不可靠的。

## 总结
Prompt 定义：由 Instruction 和 Question 组成，分别通过 system 和 user 角色传递。
Prompt 结构：包括上下文、目标、响应及示例四要素，其中示例尤为重要。

实战案例：通过三个版本的 Prompt 优化过程，展示了如何逐步改进初始 Prompt 达到预期效果。
最终方案：第三版 Prompt 成功提取出符合要求的关键词，并采用固定格式输出，便于后续处理。

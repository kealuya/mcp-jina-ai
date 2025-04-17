# MCP-Jina-AI

[English](#english) | [中文](#中文)

## English

### Introduction
MCP-Jina-AI is a powerful MCP (Model Control Protocol) tool that integrates Jina.ai's reader service to convert web content into large language model (LLM) friendly formats. It provides a simple and efficient way to transform web pages into clean, structured text that can be easily processed by AI models.

### Features
- Easy URL content extraction
- Automatic conversion to LLM-friendly format
- Clean and structured output
- Fast and reliable processing
- Built with Go for high performance

### Installation
```bash
# Install directly using go install
go install github.com/kealuya/mcp-jina-ai@latest
```

Make sure you have Go installed and your `GOPATH` is properly set up in your system's PATH.

### Usage
MCP-Jina-AI is designed to be used as a tool within the MCP ecosystem. Simply provide a URL, and it will return the content in a format optimized for large language models.

### MCP Configuration
To use mcp-jina-ai in MCP, you need to add the following configuration to your MCP configuration file:

```json
{
  "mcpServers": {
    "mcp-jina-ai": {
      "command": "mcp-jina-ai"
    }
  }
}
```

Configuration instructions:
- Add the above configuration to your MCP configuration file
- Make sure mcp-jina-ai is installed via `go install`
- Restart the MCP service to apply the configuration

### About Jina.ai Reader Service
Jina.ai's reader service (r.jina.ai) is a free service that helps extract the main content from web pages. To use it, simply add 'r.jina.ai/' before any URL you want to process.

Example:
- Original URL: `https://example.com/article`
- Jina.ai URL: `https://r.jina.ai/https://example.com/article`

## 中文

### 简介
MCP-Jina-AI 是一个强大的 MCP（模型控制协议）工具，集成了 Jina.ai 的阅读器服务，可以将网页内容转换为大语言模型（LLM）友好的格式。它提供了一种简单高效的方式，将网页转换为结构化的文本，便于 AI 模型处理。

### 特性
- 简单的 URL 内容提取
- 自动转换为 LLM 友好格式
- 清晰的结构化输出
- 快速可靠的处理
- 使用 Go 语言构建，性能优异

### 安装
```bash
# 使用 go install 直接安装
go install github.com/kealuya/mcp-jina-ai@latest
```

请确保您已安装 Go 并且在系统的 PATH 中正确设置了 `GOPATH`。

### 使用方法
MCP-Jina-AI 被设计为在 MCP 生态系统中使用的工具。只需提供一个 URL，它就会返回经过优化的、适合大语言模型使用的内容格式。

### MCP配置
要在MCP中使用mcp-jina-ai，您需要在MCP的配置文件中添加以下配置：

```json
{
  "mcpServers": {
    "mcp-jina-ai": {
      "command": "mcp-jina-ai"
    }
  }
}
```

配置说明：
- 将上述配置添加到您的MCP配置文件中
- 确保已通过`go install`命令安装了mcp-jina-ai
- 配置完成后，重启MCP服务使配置生效

### 关于 Jina.ai 阅读器服务
Jina.ai 的阅读器服务（r.jina.ai）是一个免费服务，可以帮助提取网页的主要内容。使用方法很简单，只需在任何想要处理的 URL 前面添加 'r.jina.ai/' 即可。

示例：
- 原始 URL：`https://example.com/article`
- Jina.ai URL：`https://r.jina.ai/https://example.com/article`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。




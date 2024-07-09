# Go lsp

This is an intro for me using the lsp protocal

My nvim "golsp.lua"

```lau
return {
  "golsp",
  name = "golsp",
  dev = true,
  config = function()
    local client = vim.lsp.start_client({
      name = "golsp",
      cmd = { "/home/drama321/coding/golang-lsp/main" },
    })

    if not client then
      vim.notify("custom golsp not working")
      return
    end

    vim.api.nvim_create_autocmd("FileType", {
      pattern = "markdown",
      callback = function()
        vim.lsp.buf_attach_client(0, client)
      end,
    })
  end,
}
```

I use lazynvim

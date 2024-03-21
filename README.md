# lsp-go

## installation and neo vim setup

Add the following code to your neovim configuration to start the lsp server when 
you open a markdown file. 

```lua
local client = vim.lsp.start_client {
  name = "brewinski:lsp",
  cmd = { "/Users/chris.brewin/Documents/github/lsp-go/main" },
  on_attach = on_attach,
  capabilities = capabilities,
}

if not client then
  vim.notify "Failed to start custom lsp server"
  return
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = "markdown",
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end,
})
```

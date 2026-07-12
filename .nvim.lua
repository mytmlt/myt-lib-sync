local dap = require("dap")

-- Standard Go DAP configurations
dap.configurations.go = {
	{
		type = "go",
		name = "Debug App",
		request = "launch",
		-- Point this to your main package directory (e.g., "cmd/myapp" or ".")
		program = "./cmd/youtubedr",
	},
	{
		type = "go",
		name = "Debug App with Args",
		request = "launch",
		program = "./cmd/youtubedr",
		-- If your CLI application requires arguments (like a command name)
		args = function()
			local args_string = vim.fn.input("Arguments: ")
			return vim.split(args_string, " +")
		end,
	},
}

import { McpServer, ResourceTemplate } from "@modelcontextprotocol/sdk/server/mcp.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import { z } from "zod";

import { getUsers, getUsersByName, getUserById, addUser } from "./users-tool.js";


async function main() {
  try {
    const server = new McpServer({
      name: "user-api-server",
      version: "1.0.0"
    });
  
    server.registerTool("getUsers",
      {
        title: "Retrieve all users tool",
        description: "Retrieve all users from the user API",
        inputSchema: { }
      },
      async () => {
        const users = await getUsers();
        return {
          content: [{ type: "text", text: JSON.stringify(users) }]
        }
      }
    );
  
    server.registerTool("getUsersByName",
      {
        title: "Retrieve users by name tool",
        description: "Retrieve users by name from the user API",
        inputSchema: { name: z.string() }
      },
      async ({ name }) => {
        const users = await getUsersByName(name);
        return {
          content: [{ type: "text", text: JSON.stringify(users) }]
        }
      }
    );
  
    server.registerTool("getUserById",
      {
        title: "Retrieve user by id tool",
        description: "Retrieve user by id from the user API",
        inputSchema: { id: z.string() }
      },
      async ({ id }) => {
        const user = await getUserById(id);
        return {
          content: [{ type: "text", text: JSON.stringify(user) }]
        }
      }
    );
  
    server.registerTool("registerUser",
      {
        title: "Register a new user tool",
        description: "Register a new user in the user API",
        inputSchema: { name: z.string(), email: z.string() }
      },
      async ({ name, email }) => {
        const user = await addUser({ name, email });
        return {
          content: [{ type: "text", text: `Usuário criado com sucesso: ${JSON.stringify(user)}` }]
        }
      }
    );
  
    const transport = new StdioServerTransport();
    await server.connect(transport);
  } catch (error) {
    console.error('Error initializing MCP server', error);
  }
}

main().catch((error) => {
  console.error('Error starting MCP server', error);
});
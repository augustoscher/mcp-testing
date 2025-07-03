import { User } from "./types";

const BASE_URL = "http://localhost:8080";

export async function getUsers(): Promise<User[]> {
  try {
    const response = await fetch(`${BASE_URL}/users`);
    const { data }  = await response.json();
    return data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function getUserById(id: string): Promise<User> {
  try {
    const response = await fetch(`${BASE_URL}/users/${id}`);
    const { data } = await response.json();
    return data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function getUsersByName(name: string): Promise<User[]> {
  try {
    const response = await fetch(`${BASE_URL}/users?name=${encodeURIComponent(name)}`);
    const { data } = await response.json();
    return data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function addUser(user: Omit<User, "id">): Promise<User> {
  try {
    const response = await fetch(`${BASE_URL}/users`, {
      method: "POST",
      body: JSON.stringify(user),
    });
    const { data } = await response.json();
    return data;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

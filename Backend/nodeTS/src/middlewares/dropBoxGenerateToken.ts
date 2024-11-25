import { DropboxAuth } from "dropbox";
import express, { Request, Response } from "express";
import fs from "fs";
import path from "path";
import { TokenResponse } from "../interface/tokenResponse";
import axios from "axios";

const APP_KEY = `${process.env.DROPBOX_CLIENT_ID}` || "";
const APP_SECRET = `${process.env.DROPBOX_CLIENT_SECRET_ID}` || "";
const REDIRECT_URI = "http://localhost:8000/api/v1/callback";
let currentAccessToken = `${process.env.DROPBOX_REFRESH_TOKEN}`;
let tokenExpiryTime = process.env.DROPBOX_TOKEN_EXPIRED;

const dbxAuth = new DropboxAuth({
  clientId: APP_KEY,
  clientSecret: APP_SECRET,
});

const envPath = path.resolve("./.env");

function saveTokenToEnv(key: string, value: string) {
  const envContent = fs.readFileSync(envPath, "utf8");
  const updatedContent = envContent
    .split("\n")
    .filter((line) => !line.startsWith(`${key}=`))
    .join("\n");

  fs.writeFileSync(envPath, updatedContent, "utf8");
  fs.appendFileSync(envPath, `\n${key}="${value}"`);
}

export async function getAuthUrl(req: Request, res: Response) {
  try {
    const authUrl = await dbxAuth.getAuthenticationUrl(
      REDIRECT_URI,
      undefined,
      "code",
      "offline"
    );
    res.redirect(String(authUrl));
  } catch (error) {
    console.error("Error generating auth URL:", error);
    res.status(500).send("Failed to generate auth URL");
  }
}

export async function AuthDropBox(req: Request, res: Response) {
  const { code } = req.query;

  if (!code) {
    res.status(400).send("Authorization code is missing");
    return;
  }

  try {
    const tokenResponse = await dbxAuth.getAccessTokenFromCode(
      REDIRECT_URI,
      code as string
    );

    const { access_token, refresh_token } =
      tokenResponse.result as TokenResponse;

    saveTokenToEnv("DROPBOX_ACCESS_TOKEN", access_token);
    saveTokenToEnv("DROPBOX_REFRESH_TOKEN", refresh_token || "");

    res.send("Access token and refresh token retrieved!");
  } catch (error) {
    console.error("Error getting access token:", error);
    res.status(500).send("Failed to get access token.");
  }
}

export async function refreshToken(): Promise<string> {
  if (currentAccessToken && tokenExpiryTime && Date.now() < parseInt(tokenExpiryTime)) {
    console.log("Access token is still valid.");
    return currentAccessToken;
  }

  const refreshToken = process.env.DROPBOX_REFRESH_TOKEN;
  if (!refreshToken) {
    throw new Error("Refresh token is missing");
  }

  const params = new URLSearchParams();
  params.append("grant_type", "refresh_token");
  params.append("refresh_token", refreshToken);
  params.append("client_id", APP_KEY);
  params.append("client_secret", APP_SECRET);

  try {
    const tokenResponse = await axios.post(
      "https://api.dropboxapi.com/oauth2/token",
      params
    );

    if (!tokenResponse || !tokenResponse.data) {
      throw new Error("Failed to get token response");
    }

    const { access_token, expires_in } = tokenResponse.data as TokenResponse;
    console.log("Access Token:", access_token, "Expires In:", expires_in);

    if (!access_token || !expires_in) {
      throw new Error("Missing access token or expiry time");
    }

    currentAccessToken = access_token;
    let newExpired = Date.now() + expires_in * 1000
    tokenExpiryTime = newExpired.toString()

    saveTokenToEnv("DROPBOX_ACCESS_TOKEN", access_token);
    saveTokenToEnv("DROPBOX_TOKEN_EXPIRED", tokenExpiryTime);

    return access_token;
  } catch (error: any) {
    console.error("Error refreshing access token:", error.message);
    console.error("Error stack:", error.stack);
    throw error;
  }
}
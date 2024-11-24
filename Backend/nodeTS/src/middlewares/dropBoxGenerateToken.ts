import { Dropbox, DropboxAuth } from "dropbox";
import express, { Request, Response } from "express";
import fs from "fs";
import path from "path";

const APP_KEY = `${process.env.DROPBOX_CLIENT_ID}`;
const APP_SECRET = `${process.env.DROPBOX_CLIENT_SECRET_ID}`;
const REDIRECT_URI = "http://localhost:8000/api/v1/callback";

const dbxAuth = new DropboxAuth({
  clientId: APP_KEY,
  clientSecret: APP_SECRET,
});

export async function AuthDropBox(req: Request, res: Response) {
  const { code } = req.query;

  if (!code) {
    res.status(400).send("Authorization code is missing");
    return;
  }

  try {
    const tokenResponse: any = await dbxAuth.getAccessTokenFromCode(
      REDIRECT_URI,
      code as string
    );

    const accessToken = tokenResponse.result.access_token;

    const envPath = path.resolve('./.env');
    const envContent = fs.readFileSync(envPath, "utf8");
    const updatedContent = envContent
      .split("\n")
      .filter((line) => !line.startsWith(`DROPBOX_ACCESS_TOKEN=`))
      .join("\n");

    fs.writeFileSync(envPath, updatedContent, "utf8");

    fs.appendFileSync(envPath, `\nDROPBOX_ACCESS_TOKEN="${accessToken}"`);

    res.send("Access token retrieved! Check your console.");
  } catch (error) {
    console.error("Error getting access token:", error);
    res.status(500).send("Failed to get access token.");
  }
}

export async function getAuthUrl(req: Request, res: Response) {
    try {
        const authUrl = await dbxAuth.getAuthenticationUrl(REDIRECT_URI, undefined, "code");
        res.redirect(String(authUrl));
    } catch (error) {
        res.status(500).send("Failed to generate auth URL");
    }
}

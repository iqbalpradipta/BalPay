import { Dropbox } from "dropbox";
import express from "express";
import fs from "fs";

let dropBoxAccessToken = `${process.env.DROPBOX_ACCESS_TOKEN}`;
let dbx = new Dropbox({ accessToken: dropBoxAccessToken });

export async function UploadToDropbox(
  localFilePath: string,
  dropboxPath: string
): Promise<void> {
  try {
    const fileContent = fs.readFileSync(localFilePath);
    await dbx.filesUpload({
      path: dropboxPath,
      contents: fileContent,
      mode: { ".tag": "overwrite" },
    });
  } catch (error) {
    console.error(error);
    throw error;
  } finally {
    fs.unlinkSync(localFilePath);
  }
}

export const getDropboxSharedLink = async (
  dropboxPath: string
): Promise<string> => {
  try {
    const response = await dbx.sharingCreateSharedLinkWithSettings({
      path: dropboxPath,
    });
    const modifiedLink = response.result.url.replace("&dl=0", "&raw=1");

    return modifiedLink;
  } catch (error) {
    throw error;
  }
};

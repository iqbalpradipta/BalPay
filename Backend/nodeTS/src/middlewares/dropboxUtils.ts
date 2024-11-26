import { Dropbox } from "dropbox";
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

export async function UploadToDropboxArray(
  localFilePaths: string[],
  dropboxPaths: string[]
): Promise<void> {
  if (localFilePaths.length !== dropboxPaths.length) {
    throw new Error(
      "The number of local file paths must match the number of Dropbox paths."
    );
  }

  try {
    for (let i = 0; i < localFilePaths.length; i++) {
      const fileContent = fs.readFileSync(localFilePaths[i]);

      await dbx.filesUpload({
        path: dropboxPaths[i],
        contents: fileContent,
        mode: { ".tag": "overwrite" },
      });

      console.log(`File uploaded to Dropbox: ${dropboxPaths[i]}`);

      fs.unlinkSync(localFilePaths[i]);
    }
  } catch (error) {
    console.error(error);
    throw error;
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

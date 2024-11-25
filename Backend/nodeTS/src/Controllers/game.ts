import { Request, Response } from "express";
import GameServices from "../Services/game";
import { error } from "console";
import {
  getDropboxSharedLink,
  UploadToDropbox,
} from "../middlewares/dropboxUtils";

export default new (class GameControllers {
  async CreateGame(req: Request, res: Response) {
    try {
      const pathFile = `/uploads/${req.file?.filename}`;
      const data = {
        name: req.body.name,
        description: req.body.description,
        icon: req.file?.path as string,
        types: req.body.types,
      };

      if (!data.icon) {
        res.status(400).json({ messages: "image/icon not detected" });
      } else {
        await UploadToDropbox(data.icon, pathFile);
        const getLinkFromDbx = await getDropboxSharedLink(pathFile);
        data.icon = getLinkFromDbx;
      }

      const response = await GameServices.CreateGame(data);

      res.status(200).json(response);
    } catch (error) {
      res.status(500).json(error);
    }
  }

  async GetGames(req: Request, res: Response) {
    try {
      const response = await GameServices.GetGames();

      res.status(200).json(response);
    } catch (error) {
      res.status(500).json(error);
    }
  }

  async GetGameById(req: Request, res: Response) {
    try {
      const id = parseInt(req.params.id);
      const response = await GameServices.GetGameById(id);

      if (response.status === "Failed") {
        res.status(400).json(response);
      } else {
        res.status(200).json(response);
      }
    } catch (error) {
      res.status(500).json(error);
    }
  }

  async UpdateGame(req: Request, res: Response) {
    try {
      const id = parseInt(req.params.id);
      const data = req.body;

      if (req.file) {
        const pathFile = `/uploads/${req.file?.filename}`;
        data.icon = req.file?.path as string;

        if (data.icon) {
          await UploadToDropbox(data.icon, pathFile);
          const getLinkFromDbx = await getDropboxSharedLink(pathFile);
          data.icon = getLinkFromDbx;
        }
      }

      const response = await GameServices.UpdateGame(id, data);
      if (response.status === "Failed") {
        res.status(400).json(response);
        return;
      } else {
        res.status(200).json(response);
        return;
      }
    } catch (error) {
      res.status(500).json(error);
    }
  }

  async DeleteGame(req: Request, res: Response) {
    try {
      const id = parseInt(req.params.id);
      const response = await GameServices.DeleteGame(id);

      if (response.status === "Failed") {
        res.status(400).json(response);
      } else {
        res.status(200).json(response);
      }
    } catch (error) {
      res.status(500).json(error);
    }
  }
})();

import { Request, Response } from "express";
import gameProductServices from "../Services/gameProduct";
import {
  getDropboxSharedLink,
  UploadToDropbox,
  UploadToDropboxArray,
} from "../middlewares/dropboxUtils";
import { IGameProductData } from "../interface/gameProduct";

export default new (class GameProductControllers {
  async CreateGameProduct(req: Request, res: Response) {
    try {
      const files = req.files as Express.Multer.File[];
      if (!files || files.length === 0) {
        res.status(400).json({ messages: "No images detected" });
        return;
      }

      const filePaths = files.map((file) => file.path);
      const dropboxPaths = files.map((file) => `/uploads/${file.filename}`);

      const data: IGameProductData = {
        games: { connect: { id: parseInt(req.body.games) } },
        name: req.body.name,
        price: parseFloat(req.body.price),
        image: [],
        type: req.body.type,
      };

      await UploadToDropboxArray(filePaths, dropboxPaths);

      const dropboxLinks = [];
      for (const dropboxPath of dropboxPaths) {
        const link = await getDropboxSharedLink(dropboxPath);
        dropboxLinks.push(link);
      }
      data.image = dropboxLinks;

      const response = await gameProductServices.CreateGameProduct(data);

      if (response.status == "Failed") {
        res.status(400).json(response);
      } else {
        res.status(200).json(response);
      }
    } catch (error) {
      res.status(500).json(error);
    }
  }

  async GetGameProduct(req: Request, res: Response) {
    try {
      const response = await gameProductServices.GetGameProduct();

      if (response.status == "Failed") {
        res.status(400).json(response);
      } else {
        res.status(200).json(response);
      }
    } catch (error) {
      res.status(500).json(error);
    }
  }

  async GetGameProductById(req: Request, res: Response) {
    try {
      const id = parseInt(req.params.id);
      const response = await gameProductServices.GetGameProductById(id);

      if (response.status == "Failed") {
        res.status(400).json(response);
      } else {
        res.status(200).json(response);
      }
    } catch (error) {
      res.status(500).json(error);
    }
  }

  async UpdateGameProduct(req: Request, res: Response) {
    try {
      const id = parseInt(req.params.id)
      const data: IGameProductData = {
        games: { connect: { id: parseInt(req.body.games) } },
        name: req.body.name,
        price: parseFloat(req.body.price),
        image: [],
        type: req.body.type,
      };

      if (req.file) {
        if (!data.image) {
          res.status(400).json({ messages: "image/icon not detected" });
        } else {
          const files = req.files as Express.Multer.File[];
          if (!files || files.length === 0) {
            res.status(400).json({ messages: "No images detected" });
            return;
          }

          const filePaths = files.map((file) => file.path);
          const dropboxPaths = files.map((file) => `/uploads/${file.filename}`);

          await UploadToDropboxArray(filePaths, dropboxPaths);

          const dropboxLinks = [];
          for (const dropboxPath of dropboxPaths) {
            const link = await getDropboxSharedLink(dropboxPath);
            dropboxLinks.push(link);
          }
          data.image = dropboxLinks;
        }
      }
      const response = await gameProductServices.UpdateGameProduct(id, data);

      if (response.status == "Failed") {
        res.status(400).json(response);
      } else {
        res.status(200).json(response);
      }
    } catch (error) {
      res.status(500).json(error);
    }
  }

  async DeleteGameProduct(req: Request, res: Response) {
    try {
      const id = parseInt(req.params.id);
      const response = await gameProductServices.DeleteGameProduct(id);

      if (response.status == "Failed") {
        res.status(400).json(response);
      } else {
        res.status(200).json(response);
      }
    } catch (error) {
      res.status(500).json(error);
    }
  }
})();

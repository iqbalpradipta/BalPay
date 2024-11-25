import { Request, Response } from "express";
import bcrypt from "bcrypt";
import UserServices from "../Services/user";
import {
  UploadToDropbox,
  getDropboxSharedLink,
} from "../middlewares/dropboxUtils";

export default new (class UserControllers {
  async GetUsers(req: Request, res: Response) {
    try {
      const response = await UserServices.GetUsers();

      res.status(200).json(response);
    } catch (error) {
      res.status(500).json(error);
    }
  }

  async GetUsersById(req: Request, res: Response) {
    try {
      const id = req.params.id;
      const response = await UserServices.GetUserById(parseInt(id));

      if (response.status === "Failed") {
        res.status(404).json(response);
      } else {
        res.status(200).json(response);
      }
    } catch (error) {
      res.status(500).json(error);
    }
  }

  async UpdateUsers(req: Request, res: Response): Promise<void> {
    try {
      const id = (req as any).user.id;
      const pathFile = `/uploads/${req.file?.filename}`;
      const data = {
        name: req.body.name,
        email: req.body.email,
        password: req.body.password,
        phoneNumber: req.body.phoneNumber,
        photoProfile: req.file?.path as string,
      };

      if (req.file) {
        if (!data.photoProfile) {
          res.status(400).json({ messages: "Image not detected" });
          return;
        } else {
          await UploadToDropbox(data.photoProfile, pathFile);
          const getLinkFromDbx = await getDropboxSharedLink(pathFile);
          data.photoProfile = getLinkFromDbx;
        }
      }

      if (data.password) {
        data.password = bcrypt.hashSync(req.body.password, 10);
      }

      const response = await UserServices.UpdateUser(id, data);

      if (response.status === "Failed") {
        res.status(404).json(response);
      } else {
        res.status(200).json(response);
      }
    } catch (error) {
      res.status(500).json(error);
    }
  }

  async DeleteUsers(req: Request, res: Response) {
    try {
      const id = (req as any).user.id;

      const response = await UserServices.DeleteUsers(id);

      if (response.status === "Failed") {
        res.status(404).json(response);
      } else {
        res.status(200).json(response);
      }
    } catch (error) {
      res.status(500).json(error);
    }
  }
})();

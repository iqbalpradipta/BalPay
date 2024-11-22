import { Request, Response } from "express";
import bcrypt from "bcrypt";
import AuthService from "../Services/auth";

export default new (class AuthController {
  async Register(req: Request, res: Response) {
    try {
      const data = req.body;
      data.password = bcrypt.hashSync(req.body.password, 10);
      data.photoProfile = 'https://static.vecteezy.com/system/resources/previews/036/280/650/non_2x/default-avatar-profile-icon-social-media-user-image-gray-avatar-icon-blank-profile-silhouette-illustration-vector.jpg'
      const response = await AuthService.Register(data);

      if (response.status === "Failed") {
        res.status(400).json(response);
      } else {
        res.status(200).json(response);
      }
    } catch (error) {
      res.status(500).json(error);
    }
  }

  async RegisterAdmin(req: Request, res: Response) {
    try {
      const data = req.body;
      data.password = bcrypt.hashSync(req.body.password, 10);
      data.role = "admin";
      data.photoProfile = 'https://i.pinimg.com/736x/fc/09/71/fc09710e83d5a5391ec237e0915703f2.jpg'
      const response = await AuthService.RegisterAdmin(data);

      if (response.status === "Failed") {
        res.status(400).json(response);
      } else {
        res.status(200).json(response);
      }
    } catch (error) {
      res.status(500).json(error);
    }
  }

  async Login(req: Request, res: Response) {
    try {
      const { email, password } = req.body;
      const response = await AuthService.Login(email, password);

      if (response.status === "Failed") {
        res.status(401).json(response);
      } else {
        res.status(200).json(response);
      }
    } catch (error) {
      res.status(500).json(error);
    }
  }
})();

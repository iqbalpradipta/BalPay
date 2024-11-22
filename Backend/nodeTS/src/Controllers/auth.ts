import { Request, Response } from "express";
import bcrypt from "bcrypt";
import AuthService from "../Services/auth";

export default new (class AuthController {
  async Register(req: Request, res: Response) {
    try {
      const data = req.body;
      data.password = bcrypt.hashSync(req.body.password, 10);
      const response = await AuthService.Register(data);

      if (response.status === "Failed") {
        res.status(401).json(response);
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
      const response = await AuthService.RegisterAdmin(data);

      if (response.status === "Failed") {
        res.status(401).json(response);
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

import { Prisma, PrismaClient } from "@prisma/client";
import jwt from "jsonwebtoken";
import bcrypt from "bcrypt";

const prisma = new PrismaClient();

export default new (class AuthServices {
  async Register(data: Prisma.UsersCreateInput) {
    try {
      const findEmail = await prisma.users.findFirst({
        where: {
          email: data.email,
        },
      });

      if (findEmail?.email === data.email) {
        return {
          status: "Failed",
          messages: "Email has been used! please use another email",
        };
      } else {
        const user = await prisma.users.create({ data });
        return {
          status: "Success",
          messages: "Register Success",
          data: user,
        };
      }
    } catch (error) {
      throw error;
    }
  }

  async RegisterAdmin(data: Prisma.UsersCreateInput) {
    try {
      const findEmail = await prisma.users.findFirst({
        where: {
          email: data.email,
        },
      });

      if (findEmail?.email === data.email) {
        return {
          status: "Failed",
          messages: "Email has been used! please use another email",
        };
      } else {
        const user = await prisma.users.create({ data });
        return {
          status: "Success",
          messages: "Register Success",
          data: user,
        };
      }
    } catch (error) {
      throw error;
    }
  }

  async Login(email: string, password: string) {
    try {
      const user = await prisma.users.findFirst({
        where: {
          email,
        },
      });

      const payload = {
        id: user?.id,
        name: user?.name,
        email: user?.email,
        phoneNumber: user?.phoneNumber,
        role: user?.role
      };

      const passwordCompare = bcrypt.compareSync(
        password,
        user?.password as string
      );

      if (user?.email !== email && !passwordCompare) {
        return {
          status: "Failed",
          messages: "email/password is wrong !",
        };
      } else {
        const token = jwt.sign(payload, `${process.env.SECRET_KEY_JWT}`, {
          expiresIn: "24h",
        });

        return {
          status: "Success",
          messages: "Login Success",
          token,
        };
      }
    } catch (error) {
      throw error;
    }
  }
})();

import { NextFunction, Request, Response } from "express";
import jwt, { JwtPayload } from "jsonwebtoken";

export default new class JWTAuth {
    Authentication(roles: string[]) {
        return (req: Request, res: Response, next: NextFunction): void => {
            const token = req.headers['authorization']?.split(' ')[1];

            if (!token) {
                res.status(401).json({ messages: 'No Token Provided' });
                return;
            }

            try {
                const decode = jwt.verify(token, `${process.env.SECRET_KEY_JWT}`) as JwtPayload;

                if (!roles.includes(decode.role)) {
                    res.status(403).json({ messages: 'Forbidden: Insufficient Role' });
                    return;
                }

                (req as any).user = decode;
                next();
            } catch (error) {
                res.status(401).json({ messages: 'Invalid Token' });
                return;
            }
        }
    }
}
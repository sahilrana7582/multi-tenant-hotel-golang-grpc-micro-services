import { z } from "zod";




// Props
export interface LoginFormProps {
  images: string[];
}

export const loginSchema = z.object({
  email: z.email({ message: "Invalid email address" }),
  password: z.string().min(4, { message: "Password must be at least 4 characters" })
});

export type LoginFormValues = z.infer<typeof loginSchema>;

import bcrypt from 'bcryptjs';

const hashPassword = async (password: string): Promise<string> => {
    // Generate a salt (you can adjust the salt rounds for more security)
    const saltRounds = 10;
    const salt = await bcrypt.genSalt(saltRounds);

    // Hash the password with the generated salt
    const hashedPassword = await bcrypt.hash(password, salt);

    return hashedPassword;
};
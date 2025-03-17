
export const rules = {
  required: (v) => !!v || "Harus diisi",
  email: (v) => /.+@.+\..+/.test(v) || "E-mail harus valid",
  passwordMinLength: (length) => (v) =>
    (v && v.length >= length) || `Password harus terdiri dari ${length} characters`,
}
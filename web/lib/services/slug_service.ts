export function splitSlug(app: string) {
  const slug = app.split("-");
  return {
    id: slug[0],
    slug: slug[1],
  };
}

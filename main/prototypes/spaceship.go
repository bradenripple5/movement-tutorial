components {
  id: "spaceship"
  component: "/main/spaceship.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"spaceship\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/sprites.atlas\"\n"
  "}\n"
  ""
}

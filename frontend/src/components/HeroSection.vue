<template>
  <section class="section">
    <div class="section__hero hero container">
      <h1 class="hero__title">Не только людям нужен дом</h1>
      <p class="hero__subtitle">
        Подарите шанс маленькому щенку с очень большим сердцем. Он или она будет любить Вас больше
        всех на свете, вот увидите!
      </p>
      <router-link to="/animals" class="hero__button button button--header">
        Подобрать друга
      </router-link>
    </div>
  </section>
</template>

<script setup></script>

<style lang="scss" scoped>
.hero {
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: rem(40);
  height: rem(882);
  background-image: url('../../images/start-screen-puppy.png');
  background-repeat: no-repeat;
  background-position: 100% 100%;
  background-size: 50%;
  height: 100vh;

  @include tablet {
    margin-top: rem(-30);
    background-position: 100% 100%;
  }

  @include mobile-s {
    background-size: 50%;
  }

  &__title {
    font-size: rem(44);
    font-weight: 400;
    line-height: 130%;
    max-width: rem(373);

    @include mobile {
      font-size: rem(34);
    }
  }

  &__subtitle {
    font-size: rem(15);
    color: var(--color-light);
    line-height: 160%;
    max-width: rem(351);
  }
}
</style>

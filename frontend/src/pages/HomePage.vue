<script setup>
import Header from '@/components/Header.vue'
import HeroSection from '@/components/HeroSection.vue'
import AboutSection from '@/components/AboutSection.vue'
import HowItWorksSection from '@/components/HowItWorksSection.vue'
import UsabilitySection from '@/components/UsabilitySection.vue'
import AnimalsSlider from '@/components/AnimalsSlider.vue'
import Footer from '@/components/Footer.vue'
</script>

<template>
  <Header />
  <HeroSection />
  <AboutSection />
  <HowItWorksSection />
  <UsabilitySection />
  <AnimalsSlider />
  <Footer />
</template>

<style lang="scss">


.section {
  overflow-x: hidden;
  &:nth-child(2) {
    background-color: #262425;
    background-image:
      url('/images/noise_transparent.png'),
      linear-gradient(
        243deg,
        rgba(38, 36, 37, 1) 0%,
        rgba(106, 81, 63, 1) 65%,
        rgba(55, 46, 40, 1) 100%
      );
    background-position: 0 0;
    background-size:
      110px 110px,
      auto;
    background-repeat: repeat, no-repeat;
    margin-top: rem(-90);

    @include mobile {
      margin-top: rem(-130);
    }
  }

  &:nth-child(4) {
    background-color: #f6f6f6;
  }

  &:nth-child(5) {
    background-color: #262425;
    background-image:
      url('/images/noise_transparent.png'),
      linear-gradient(
        129deg,
        rgba(38, 36, 37, 1) 0%,
        rgba(106, 81, 63, 1) 65%,
        rgba(55, 46, 40, 1) 100%
      );
    background-position: 0 0;
    background-size:
      110px 110px,
      auto;
    background-repeat: repeat, no-repeat;
  }

  &:nth-child(6) {
    background-color: var(--color-gray-2);
  }

  &__about,
  &__howitworks {
    padding-top: rem(80);
    padding-bottom: rem(100);

    @include tablet {
      padding-top: rem(40);
      padding-bottom: rem(50);
    }
  }

  &__usability {
    padding-top: rem(80);
    padding-bottom: rem(320);
    position: relative;

    @include mobile {
      padding-top: rem(60);
    }

    &::before {
      content: '';
      position: absolute;
      background-image: url('/images/decor_usability.png');
      width: 100%;
      height: 100%;
      overflow: hidden;
      background-repeat: no-repeat;
      z-index: 0;
    }
  }

  &__animals {
    padding-top: rem(80);
    padding-bottom: rem(100);

    @include mobile {
      padding-top: rem(40);
      padding-bottom: rem(50);
    }
  }

  &--transparent-bg {
    background: var(--color-gray-2) !important;
    margin-top: 0 !important;
  }

  &__findhome {
    padding-bottom: rem(54);
  }
}

.swiper-wrapper {
  padding-block: rem(40) !important;
}
</style>

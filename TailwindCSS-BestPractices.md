# TailwindCSS v4.0+ 实际项目最佳实践指南

## 🎯 核心理念

TailwindCSS是一个**原子化CSS框架**，采用"utility-first"设计理念，通过组合小而精的CSS类来构建界面。

## 🚀 Tailwind CSS v4.0 重大变化

### **⚡ 性能提升**
- **全量构建速度提升 3.5x+**
- **增量构建速度提升 8x+**
- **微秒级构建时间**（无新 CSS 时）

### **� 现代 Web 技术**
- 原生 CSS 层级（`@layer`）
- 注册自定义属性（`@property`）
- `color-mix()` 函数支持
- 逻辑属性自动支持

## �🏗️ 项目安装与配置

### 1. **快速安装（v4.0+）**

```bash
# 安装 TailwindCSS v4
npm install tailwindcss @tailwindcss/vite

# 或使用 PostCSS
npm install tailwindcss @tailwindcss/postcss
```

### 2. **Vite 项目配置**

```typescript
// vite.config.ts
import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
  plugins: [
    tailwindcss(),
  ],
})
```

### 3. **PostCSS 配置**

```javascript
// postcss.config.js
export default {
  plugins: ["@tailwindcss/postcss"],
}
```

### 4. **CSS 文件配置（v4.0 新语法）**

```css
/* styles/main.css */
@import "tailwindcss";

/* 🎨 主题变量配置 - v4.0 新特性 */
@theme {
  /* 字体配置 */
  --font-display: "Satoshi", sans-serif;
  --font-body: "Inter", system-ui, sans-serif;
  
  /* 色彩系统 - 使用 OKLCH */
  --color-primary-50: oklch(0.98 0.02 260);
  --color-primary-100: oklch(0.95 0.05 260);
  --color-primary-500: oklch(0.65 0.2 260);
  --color-primary-900: oklch(0.3 0.15 260);
  
  /* 品牌色彩 */
  --color-brand: oklch(0.65 0.2 260);
  --color-accent: oklch(0.75 0.15 120);
  
  /* 间距系统 */
  --spacing: 0.25rem; /* 基础间距单位 */
  
  /* 断点配置 */
  --breakpoint-3xl: 1920px;
  --breakpoint-4xl: 2560px;
  
  /* 容器查询断点 */
  --container-xs: 20rem;
  --container-sm: 24rem;
  --container-md: 28rem;
  --container-lg: 32rem;
  --container-xl: 36rem;
  
  /* 阴影系统 */
  --shadow-brutal: 8px 8px 0px 0px oklch(0 0 0);
  --shadow-soft: 0 10px 25px -5px oklch(0 0 0 / 0.1);
  
  /* 动画缓动 */
  --ease-bounce: cubic-bezier(0.68, -0.55, 0.265, 1.55);
  --ease-smooth: cubic-bezier(0.4, 0, 0.2, 1);
}

/* 🔧 自定义组件样式 */
@layer components {
  .btn {
    @apply px-4 py-2 rounded-lg font-medium transition-all duration-200;
    @apply focus:outline-none focus:ring-2 focus:ring-offset-2;
  }
  
  .btn-primary {
    @apply bg-primary-500 text-white hover:bg-primary-600;
    @apply focus:ring-primary-500;
  }
  
  .btn-secondary {
    @apply bg-gray-200 text-gray-900 hover:bg-gray-300;
    @apply dark:bg-gray-700 dark:text-gray-100 dark:hover:bg-gray-600;
  }
  
  .card {
    @apply bg-white rounded-xl shadow-lg border border-gray-200;
    @apply dark:bg-gray-800 dark:border-gray-700;
  }
  
  .input {
    @apply w-full px-3 py-2 border border-gray-300 rounded-lg;
    @apply focus:ring-2 focus:ring-primary-500 focus:border-primary-500;
    @apply dark:bg-gray-700 dark:border-gray-600 dark:text-white;
  }
}

/* 🛠️ 自定义工具类 */
@layer utilities {
  .text-balance {
    text-wrap: balance;
  }
  
  .scrollbar-none {
    scrollbar-width: none;
    -ms-overflow-style: none;
  }
  
  .scrollbar-none::-webkit-scrollbar {
    display: none;
  }
  
  /* 3D 变换支持 */
  .transform-3d {
    transform-style: preserve-3d;
  }
  
  /* 渐变工具类 */
  .gradient-mesh {
    background: 
      radial-gradient(at 40% 20%, oklch(0.7 0.15 300) 0px, transparent 50%),
      radial-gradient(at 80% 0%, oklch(0.6 0.2 260) 0px, transparent 50%),
      radial-gradient(at 0% 50%, oklch(0.65 0.18 200) 0px, transparent 50%);
  }
}

/* 🌙 暗色主题 */
@layer base {
  :root {
    --background: oklch(1 0 0);
    --foreground: oklch(0.09 0 0);
  }
  
  .dark {
    --background: oklch(0.09 0 0);
    --foreground: oklch(0.98 0 0);
  }
  
  body {
    @apply bg-[var(--background)] text-[var(--foreground)];
    @apply transition-colors duration-300;
  }
}
```

### 5. **内容检测（v4.0 自动化）**

```css
/* 自动内容检测 - 无需配置 content 数组 */
@import "tailwindcss";

/* 手动添加额外源（可选） */
@source "../node_modules/@my-company/ui-lib";
@source "./src/external/**/*.html";
```

## 💡 v4.0 新特性实践

### 7. **组件化思维**

```jsx
// ✅ v4.0 支持任意网格列数
function DynamicGrid({ columns = 12 }) {
  return (
    <div className={`grid grid-cols-${columns} gap-4`}>
      {/* 内容 */}
    </div>
  );
}

// ✅ v4.0 支持任意间距值
function DynamicSpacing() {
  return (
    <div className="mt-23 pr-17 w-45 h-128">
      {/* 任意数值都会基于 --spacing 变量计算 */}
    </div>
  );
}

// ✅ v4.0 支持动态数据属性
function StateComponent({ isActive }) {
  return (
    <div 
      data-active={isActive}
      className="opacity-50 data-active:opacity-100 transition-opacity"
    >
      动态状态样式
    </div>
  );
}
```

### 2. **容器查询（v4.0 内置支持）**

```jsx
// ✅ 容器查询 - 无需插件
function ResponsiveCard() {
  return (
    <div className="@container bg-white rounded-lg p-4">
      <div className="grid grid-cols-1 @sm:grid-cols-2 @lg:grid-cols-3 gap-4">
        <div className="aspect-square @sm:aspect-video">
          <img 
            src="/image.jpg" 
            className="w-full h-full object-cover rounded"
          />
        </div>
        <div className="@sm:col-span-1 @lg:col-span-2">
          <h3 className="text-lg @sm:text-xl @lg:text-2xl font-bold">
            响应式标题
          </h3>
          <p className="text-sm @sm:text-base text-gray-600 mt-2">
            基于容器大小自适应的内容
          </p>
        </div>
      </div>
    </div>
  );
}

// ✅ 容器查询范围
function ContainerRanges() {
  return (
    <div className="@container">
      <div className="
        bg-red-100 
        @min-md:@max-xl:bg-blue-100
        @min-xl:bg-green-100
      ">
        容器查询范围示例
      </div>
    </div>
  );
}
```

### 3. **3D 变换（v4.0 新增）**

```jsx
// ✅ 3D 变换支持
function Card3D() {
  return (
    <div className="perspective-distant">
      <div className="
        transform-3d 
        rotate-x-12 
        rotate-y-6 
        rotate-z-3
        transition-transform 
        duration-300 
        hover:rotate-x-0 
        hover:rotate-y-0
        bg-white 
        rounded-xl 
        p-6 
        shadow-xl
      ">
        <h3 className="text-xl font-bold">3D 卡片</h3>
        <p className="text-gray-600 mt-2">
          悬停时恢复到原始角度
        </p>
      </div>
    </div>
  );
}

// ✅ 3D 相册效果
function PhotoGallery3D() {
  return (
    <div className="perspective-normal flex items-center justify-center space-x-4">
      {[...Array(5)].map((_, i) => (
        <div
          key={i}
          className={`
            transform-3d
            translate-z-${i * 20}
            rotate-y-${(i - 2) * 15}
            transition-all
            duration-500
            hover:translate-z-40
            hover:rotate-y-0
          `}
        >
          <img 
            src={`/photo-${i + 1}.jpg`}
            className="w-48 h-64 object-cover rounded-lg shadow-lg"
          />
        </div>
      ))}
    </div>
  );
}
```

### 4. **增强渐变系统（v4.0）**

```jsx
// ✅ 角度渐变
function AngleGradients() {
  return (
    <div className="space-y-4">
      <div className="h-32 bg-linear-45 from-purple-500 to-pink-500 rounded-lg" />
      <div className="h-32 bg-linear-135 from-blue-500 to-teal-400 rounded-lg" />
      <div className="h-32 bg-linear-to-r/oklch from-red-500 to-yellow-400 rounded-lg" />
    </div>
  );
}

// ✅ 圆锥和径向渐变
function AdvancedGradients() {
  return (
    <div className="grid grid-cols-2 gap-4">
      {/* 圆锥渐变 */}
      <div className="
        size-32 
        rounded-full 
        bg-conic/[in_hsl_longer_hue] 
        from-red-600 
        to-red-600
      " />
      
      {/* 径向渐变 */}
      <div className="
        size-32 
        rounded-full 
        bg-radial-[at_25%_25%] 
        from-white 
        to-zinc-900 
        to-75%
      " />
      
      {/* 复杂渐变组合 */}
      <div className="
        col-span-2 
        h-32 
        bg-linear-to-br/oklch 
        from-purple-500 
        via-pink-500 
        to-orange-500
        rounded-lg
      " />
    </div>
  );
}
```

### 5. **@starting-style 过渡（v4.0）**

```jsx
// ✅ 元素首次显示动画
function PopoverWithTransition() {
  return (
    <div>
      <button 
        popovertarget="notification-popover"
        className="btn btn-primary"
      >
        显示通知
      </button>
      
      <div
        popover
        id="notification-popover"
        className="
          p-4 
          bg-white 
          rounded-lg 
          shadow-xl 
          border
          transition-all
          duration-300
          starting:opacity-0
          starting:scale-95
          starting:translate-y-2
        "
      >
        <h3 className="font-bold">通知标题</h3>
        <p className="text-gray-600 mt-1">
          这是一个使用 @starting-style 的原生过渡效果
        </p>
      </div>
    </div>
  );
}

// ✅ 列表项动画
function AnimatedList({ items }) {
  return (
    <ul className="space-y-2">
      {items.map((item, index) => (
        <li
          key={item.id}
          style={{ '--delay': `${index * 100}ms` }}
          className="
            p-3 
            bg-gray-100 
            rounded-lg
            transition-all
            duration-300
            delay-[var(--delay)]
            starting:opacity-0
            starting:translate-x-4
          "
        >
          {item.name}
        </li>
      ))}
    </ul>
  );
}
```

### 6. **not-* 变体（v4.0）**

```jsx
// ✅ 反向条件样式
function ConditionalStyling() {
  return (
    <div className="space-y-4">
      {/* 非悬停状态 */}
      <button className="
        px-4 py-2 
        bg-blue-500 
        text-white 
        rounded-lg
        not-hover:opacity-75
        transition-opacity
      ">
        悬停时完全不透明
      </button>
      
      {/* 不支持某些特性时的后备样式 */}
      <div className="
        p-4 
        bg-white 
        rounded-lg
        supports-backdrop-blur:bg-white/80
        supports-backdrop-blur:backdrop-blur-sm
        not-supports-backdrop-blur:bg-white
        not-supports-backdrop-blur:border
      ">
        智能后备样式
      </div>
      
      {/* 非触摸设备样式 */}
      <div className="
        p-4 
        bg-gray-100 
        rounded-lg
        not-hover:cursor-default
        hover:bg-gray-200
      ">
        仅在支持悬停的设备上显示悬停效果
      </div>
    </div>
  );
}
```

## 💡 代码组织最佳实践

```jsx
// ❌ 不好的做法：直接在JSX中写大量类名
function ProductCard({ product }) {
  return (
    <div className="bg-white rounded-lg shadow-lg p-6 border border-gray-200 hover:shadow-xl transition-shadow duration-300 cursor-pointer">
      <img className="w-full h-48 object-cover rounded-lg mb-4" src={product.image} />
      <h3 className="text-xl font-semibold text-gray-900 mb-2">{product.name}</h3>
      <p className="text-gray-600 text-sm mb-4">{product.description}</p>
      <div className="flex justify-between items-center">
        <span className="text-2xl font-bold text-green-600">${product.price}</span>
        <button className="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-colors">
          Add to Cart
        </button>
      </div>
    </div>
  );
}
```

```jsx
// ✅ 好的做法：使用组件抽象和样式复用
function ProductCard({ product }) {
  return (
    <Card className="hover:shadow-xl transition-shadow duration-300 cursor-pointer">
      <CardImage src={product.image} alt={product.name} />
      <CardContent>
        <CardTitle>{product.name}</CardTitle>
        <CardDescription>{product.description}</CardDescription>
        <CardFooter>
          <Price amount={product.price} />
          <Button variant="primary">Add to Cart</Button>
        </CardFooter>
      </CardContent>
    </Card>
  );
}

// 子组件定义
const Card = ({ children, className = "" }) => (
  <div className={`card ${className}`}>
    {children}
  </div>
);

const CardImage = ({ src, alt }) => (
  <img className="w-full h-48 object-cover rounded-lg mb-4" src={src} alt={alt} />
);

const Button = ({ children, variant = "primary", ...props }) => {
  const baseClasses = "px-4 py-2 rounded-lg transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2";
  const variants = {
    primary: "bg-blue-500 text-white hover:bg-blue-600 focus:ring-blue-500",
    secondary: "bg-gray-200 text-gray-900 hover:bg-gray-300 focus:ring-gray-500",
  };
  
  return (
    <button className={`${baseClasses} ${variants[variant]}`} {...props}>
      {children}
    </button>
  );
};
```

### 8. **工具函数和样式组合**

```javascript
// utils/classNames.js - v4.0 兼容
export function classNames(...classes) {
  return classes.filter(Boolean).join(' ');
}

// 使用 clsx 或 classnames 库
import clsx from 'clsx';

// ✅ v4.0 推荐的组件变体系统
function Button({ variant, size, disabled, children, ...props }) {
  return (
    <button
      className={clsx(
        // 基础样式
        'inline-flex items-center justify-center rounded-lg font-medium',
        'transition-all duration-200 transform-gpu',
        'focus:outline-none focus:ring-2 focus:ring-offset-2',
        
        // 变体样式 - 使用 CSS 变量
        {
          'bg-primary-500 text-white hover:bg-primary-600 focus:ring-primary-500': variant === 'primary',
          'bg-gray-200 text-gray-900 hover:bg-gray-300 focus:ring-gray-500': variant === 'secondary',
          'border border-gray-300 text-gray-700 hover:bg-gray-50 focus:ring-gray-500': variant === 'outline',
          'bg-red-500 text-white hover:bg-red-600 focus:ring-red-500': variant === 'danger',
        },
        
        // 尺寸样式 - v4.0 动态间距
        {
          'px-3 py-1.5 text-sm h-8': size === 'sm',
          'px-4 py-2 text-base h-10': size === 'md',
          'px-6 py-3 text-lg h-12': size === 'lg',
        },
        
        // 状态样式
        {
          'opacity-50 cursor-not-allowed': disabled,
          'hover:scale-105 active:scale-95': !disabled,
        }
      )}
      disabled={disabled}
      {...props}
    >
      {children}
    </button>
  );
}

// ✅ v4.0 主题变量引用
function ThemeAwareButton({ variant = 'primary', children, ...props }) {
  const variantStyles = {
    primary: 'bg-[var(--color-primary-500)] text-white hover:bg-[var(--color-primary-600)]',
    secondary: 'bg-[var(--color-gray-200)] text-[var(--color-gray-900)] hover:bg-[var(--color-gray-300)]',
    accent: 'bg-[var(--color-accent)] text-white hover:opacity-90',
  };

  return (
    <button
      className={`
        px-[calc(var(--spacing)*4)] 
        py-[calc(var(--spacing)*2)] 
        rounded-[var(--radius-lg)]
        transition-all 
        duration-200
        ${variantStyles[variant]}
      `}
      {...props}
    >
      {children}
    </button>
  );
}
```

## 🎨 响应式设计最佳实践

### 9. **移动优先原则**

```jsx
// ✅ v4.0 移动优先，逐步增强
function ResponsiveLayout() {
  return (
    <div className="
      grid 
      grid-cols-1 
      gap-4 
      p-4
      
      sm:grid-cols-2 
      sm:gap-6 
      sm:p-6
      
      md:grid-cols-3 
      md:gap-8 
      md:p-8
      
      lg:grid-cols-4 
      lg:gap-10 
      lg:p-10
      
      xl:max-w-7xl 
      xl:mx-auto
      
      2xl:grid-cols-6
      3xl:grid-cols-8
    ">
      {/* 内容 */}
    </div>
  );
}

// ✅ v4.0 容器查询 + 媒体查询混合
function HybridResponsive() {
  return (
    <div className="@container">
      <div className="
        grid
        grid-cols-1
        gap-2
        
        /* 容器查询 */
        @sm:grid-cols-2
        @md:grid-cols-3
        @lg:grid-cols-4
        
        /* 媒体查询 */
        sm:gap-4
        md:gap-6
        lg:gap-8
        
        /* 视口查询 */
        xl:@lg:grid-cols-6
      ">
        {/* 内容 */}
      </div>
    </div>
  );
}
```

### 10. **断点使用策略**

```css
/* v4.0 @theme 中定义断点 */
@theme {
  /* 标准断点 */
  --breakpoint-xs: 475px;
  --breakpoint-sm: 640px;
  --breakpoint-md: 768px;
  --breakpoint-lg: 1024px;
  --breakpoint-xl: 1280px;
  --breakpoint-2xl: 1536px;
  
  /* 自定义断点 */
  --breakpoint-3xl: 1920px;
  --breakpoint-4xl: 2560px;
  --breakpoint-mobile: 480px;
  --breakpoint-tablet: 768px;
  --breakpoint-laptop: 1024px;
  --breakpoint-desktop: 1440px;
  
  /* 容器查询断点 */
  --container-xs: 20rem;
  --container-sm: 24rem;
  --container-md: 28rem;
  --container-lg: 32rem;
  --container-xl: 36rem;
  --container-2xl: 42rem;
}
```

```jsx
// ✅ v4.0 使用自定义断点
function ResponsiveCard() {
  return (
    <div className="
      w-full 
      mobile:w-1/2 
      tablet:w-1/3 
      laptop:w-1/4
      desktop:w-1/5
      
      p-4 
      mobile:p-6 
      tablet:p-8
      
      text-sm 
      mobile:text-base 
      tablet:text-lg
      desktop:text-xl
    ">
      {/* 内容 */}
    </div>
  );
}
```

### 11. **自动内容检测和优化**

```css
/* v4.0 自动内容检测 - 无需配置 */
@import "tailwindcss";

/* 自动忽略：
   - .gitignore 中的文件
   - node_modules（除非明确指定）
   - 二进制文件（图片、视频等）
   - 构建产物
*/

/* 手动包含特定源 */
@source "../node_modules/@my-company/ui-lib";
@source "./src/external/**/*.html";
@source "../shared-components/**/*.{js,ts,jsx,tsx}";
```

### 12. **避免动态类名**

```jsx
// ❌ 仍需避免这样做（可能被优化掉）
function DynamicButton({ color }) {
  return (
    <button className={`bg-${color}-500 text-white`}>
      Click me
    </button>
  );
}

// ✅ v4.0 推荐做法 - 使用 CSS 变量
function DynamicButton({ color }) {
  const colorMap = {
    red: 'bg-red-500 text-white',
    green: 'bg-green-500 text-white',
    blue: 'bg-blue-500 text-white',
    // 或使用 CSS 变量
    primary: 'bg-[var(--color-primary-500)] text-white',
    secondary: 'bg-[var(--color-secondary-500)] text-white',
  };
  
  return (
    <button className={colorMap[color] || colorMap.primary}>
      Click me
    </button>
  );
}

// ✅ v4.0 更好的做法 - 利用动态值支持
function ModernDynamicButton({ colorIntensity = 500 }) {
  // v4.0 支持动态数值
  return (
    <button className={`bg-blue-${colorIntensity} text-white hover:bg-blue-${colorIntensity + 100}`}>
      动态强度按钮
    </button>
  );
}
```

## 🎭 主题和设计系统（v4.0）

### 1. **CSS 变量设计令牌**

```css
/* v4.0 @theme 语法 */
@theme {
  /* 🎨 色彩系统 - 使用 OKLCH */
  --color-brand-50: oklch(0.98 0.02 260);
  --color-brand-100: oklch(0.95 0.05 260);
  --color-brand-200: oklch(0.90 0.10 260);
  --color-brand-300: oklch(0.83 0.15 260);
  --color-brand-400: oklch(0.75 0.18 260);
  --color-brand-500: oklch(0.65 0.20 260); /* 主色 */
  --color-brand-600: oklch(0.55 0.22 260);
  --color-brand-700: oklch(0.45 0.20 260);
  --color-brand-800: oklch(0.35 0.18 260);
  --color-brand-900: oklch(0.25 0.15 260);
  --color-brand-950: oklch(0.15 0.10 260);
  
  /* 🌈 语义化色彩 */
  --color-success: oklch(0.65 0.15 142);
  --color-warning: oklch(0.75 0.15 85);
  --color-error: oklch(0.60 0.20 25);
  --color-info: oklch(0.65 0.15 230);
  
  /* 🔤 字体系统 */
  --font-sans: "Inter Variable", "Inter", system-ui, sans-serif;
  --font-serif: "Merriweather", Georgia, serif;
  --font-mono: "JetBrains Mono Variable", "Fira Code", monospace;
  --font-display: "Satoshi Variable", "Satoshi", sans-serif;
  
  /* 📏 文字大小 */
  --text-xs: 0.75rem;
  --text-xs--line-height: 1rem;
  --text-sm: 0.875rem;
  --text-sm--line-height: 1.25rem;
  --text-base: 1rem;
  --text-base--line-height: 1.5rem;
  --text-lg: 1.125rem;
  --text-lg--line-height: 1.75rem;
  --text-xl: 1.25rem;
  --text-xl--line-height: 1.75rem;
  --text-2xl: 1.5rem;
  --text-2xl--line-height: 2rem;
  --text-3xl: 1.875rem;
  --text-3xl--line-height: 2.25rem;
  
  /* 📐 间距系统 */
  --spacing: 0.25rem; /* 4px 基础单位 */
  
  /* 🔄 圆角系统 */
  --radius-xs: 0.125rem;
  --radius-sm: 0.25rem;
  --radius-md: 0.375rem;
  --radius-lg: 0.5rem;
  --radius-xl: 0.75rem;
  --radius-2xl: 1rem;
  --radius-3xl: 1.5rem;
  --radius-full: 9999px;
  
  /* 🌊 阴影系统 */
  --shadow-xs: 0 1px 2px 0 oklch(0 0 0 / 0.05);
  --shadow-sm: 0 1px 3px 0 oklch(0 0 0 / 0.1), 0 1px 2px -1px oklch(0 0 0 / 0.1);
  --shadow-md: 0 4px 6px -1px oklch(0 0 0 / 0.1), 0 2px 4px -2px oklch(0 0 0 / 0.1);
  --shadow-lg: 0 10px 15px -3px oklch(0 0 0 / 0.1), 0 4px 6px -4px oklch(0 0 0 / 0.1);
  --shadow-xl: 0 20px 25px -5px oklch(0 0 0 / 0.1), 0 8px 10px -6px oklch(0 0 0 / 0.1);
  --shadow-2xl: 0 25px 50px -12px oklch(0 0 0 / 0.25);
  
  /* ⚡ 动画缓动 */
  --ease-linear: linear;
  --ease-in: cubic-bezier(0.4, 0, 1, 1);
  --ease-out: cubic-bezier(0, 0, 0.2, 1);
  --ease-in-out: cubic-bezier(0.4, 0, 0.2, 1);
  --ease-bounce: cubic-bezier(0.68, -0.55, 0.265, 1.55);
  --ease-elastic: cubic-bezier(0.175, 0.885, 0.32, 1.275);
  
  /* 📱 断点系统 */
  --breakpoint-xs: 475px;
  --breakpoint-sm: 640px;
  --breakpoint-md: 768px;
  --breakpoint-lg: 1024px;
  --breakpoint-xl: 1280px;
  --breakpoint-2xl: 1536px;
  --breakpoint-3xl: 1920px;
  
  /* 📦 容器查询 */
  --container-xs: 20rem;
  --container-sm: 24rem;
  --container-md: 28rem;
  --container-lg: 32rem;
  --container-xl: 36rem;
  --container-2xl: 42rem;
}
```

### 2. **暗色主题支持（v4.0）**

```css
/* v4.0 现代暗色主题实现 */
@layer base {
  /* 🌞 亮色主题变量 */
  :root {
    --background: oklch(1 0 0);
    --foreground: oklch(0.09 0 0);
    --muted: oklch(0.96 0 0);
    --muted-foreground: oklch(0.45 0 0);
    --card: oklch(1 0 0);
    --card-foreground: oklch(0.09 0 0);
    --border: oklch(0.9 0 0);
    --input: oklch(0.9 0 0);
    --primary: oklch(0.24 0 0);
    --primary-foreground: oklch(0.98 0 0);
    --secondary: oklch(0.96 0 0);
    --secondary-foreground: oklch(0.09 0 0);
    --accent: oklch(0.96 0 0);
    --accent-foreground: oklch(0.09 0 0);
    --destructive: oklch(0.58 0.15 27);
    --destructive-foreground: oklch(0.98 0 0);
    --ring: oklch(0.24 0 0);
  }
  
  /* 🌙 暗色主题变量 */
  .dark {
    --background: oklch(0.04 0 0);
    --foreground: oklch(0.98 0 0);
    --muted: oklch(0.07 0 0);
    --muted-foreground: oklch(0.65 0 0);
    --card: oklch(0.04 0 0);
    --card-foreground: oklch(0.98 0 0);
    --border: oklch(0.15 0 0);
    --input: oklch(0.15 0 0);
    --primary: oklch(0.98 0 0);
    --primary-foreground: oklch(0.09 0 0);
    --secondary: oklch(0.15 0 0);
    --secondary-foreground: oklch(0.98 0 0);
    --accent: oklch(0.15 0 0);
    --accent-foreground: oklch(0.98 0 0);
    --destructive: oklch(0.62 0.18 29);
    --destructive-foreground: oklch(0.98 0 0);
    --ring: oklch(0.83 0 0);
    
    /* 暗色模式的色彩方案 */
    color-scheme: dark;
  }
  
  /* 基础样式 */
  body {
    @apply bg-[var(--background)] text-[var(--foreground)];
    @apply transition-colors duration-300;
  }
}

/* 组件样式更新 */
@layer components {
  .card {
    @apply bg-[var(--card)] text-[var(--card-foreground)];
    @apply border border-[var(--border)];
  }
  
  .btn-primary {
    @apply bg-[var(--primary)] text-[var(--primary-foreground)];
    @apply hover:bg-[var(--primary)]/90;
  }
  
  .btn-secondary {
    @apply bg-[var(--secondary)] text-[var(--secondary-foreground)];
    @apply hover:bg-[var(--secondary)]/80;
  }
  
  .input {
    @apply bg-[var(--background)] border-[var(--border)];
    @apply text-[var(--foreground)] placeholder:text-[var(--muted-foreground)];
  }
}
```

```jsx
// ✅ v4.0 主题切换组件
function ThemeToggle() {
  const [isDark, setIsDark] = useState(false);
  
  useEffect(() => {
    const root = document.documentElement;
    if (isDark) {
      root.classList.add('dark');
    } else {
      root.classList.remove('dark');
    }
  }, [isDark]);
  
  return (
    <button
      onClick={() => setIsDark(!isDark)}
      className="
        relative
        p-2 
        rounded-lg 
        bg-[var(--muted)] 
        text-[var(--muted-foreground)]
        hover:bg-[var(--accent)]
        hover:text-[var(--accent-foreground)]
        transition-all 
        duration-300
        group
      "
    >
      {/* 太阳图标 */}
      <svg
        className={`
          w-5 h-5 
          transition-all 
          duration-300 
          ${isDark ? 'rotate-90 scale-0' : 'rotate-0 scale-100'}
        `}
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
      </svg>
      
      {/* 月亮图标 */}
      <svg
        className={`
          absolute 
          top-2 
          left-2 
          w-5 h-5 
          transition-all 
          duration-300 
          ${isDark ? 'rotate-0 scale-100' : '-rotate-90 scale-0'}
        `}
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
      </svg>
    </button>
  );
}

// ✅ 系统主题检测
function SystemThemeProvider({ children }) {
  useEffect(() => {
    const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
    const handleChange = (e) => {
      document.documentElement.classList.toggle('dark', e.matches);
    };
    
    // 初始设置
    handleChange(mediaQuery);
    
    // 监听变化
    mediaQuery.addEventListener('change', handleChange);
    
    return () => mediaQuery.removeEventListener('change', handleChange);
  }, []);
  
  return children;
}
```

## 🚀 开发工具和工作流（v4.0）

### 1. **VS Code 配置**

```json
// .vscode/settings.json
{
  "tailwindCSS.includeLanguages": {
    "javascript": "javascript",
    "typescript": "typescript", 
    "html": "html",
    "vue": "html",
    "svelte": "html"
  },
  "editor.quickSuggestions": {
    "strings": true
  },
  "tailwindCSS.classAttributes": [
    "class",
    "className",
    "ngClass",
    "class:list"
  ],
  "tailwindCSS.experimental.classRegex": [
    // 支持 clsx, classnames 等库
    ["clsx\\(([^)]*)\\)", "\"([^\"]*)\""],
    ["classnames\\(([^)]*)\\)", "\"([^\"]*)\""],
    ["cva\\(([^)]*)\\)", "[\"'`]([^\"'`]*).*?[\"'`]"],
    // 支持 @theme 变量
    ["@theme\\s*{([^}]*)}", "--[\\w-]+\\s*:\\s*[^;]+"],
  ]
}
```

### 2. **ESLint 和 Prettier 配置**

```javascript
// eslint.config.js (v4.0 兼容)
import js from '@eslint/js';
import tailwind from 'eslint-plugin-tailwindcss';

export default [
  js.configs.recommended,
  {
    plugins: {
      tailwindcss: tailwind,
    },
    rules: {
      'tailwindcss/classnames-order': 'warn',
      'tailwindcss/no-custom-classname': 'warn',
      'tailwindcss/no-contradicting-classname': 'error',
      'tailwindcss/enforces-negative-arbitrary-values': 'warn',
      'tailwindcss/enforces-shorthand': 'warn',
      'tailwindcss/migration-from-tailwind-2': 'warn',
      'tailwindcss/no-arbitrary-value': 'off', // v4.0 支持更多任意值
    },
    settings: {
      tailwindcss: {
        // v4.0 CSS 配置路径
        config: './src/styles/main.css',
        // 或指定 JS 配置文件
        // config: './tailwind.config.js',
      },
    },
  },
];
```

```javascript
// prettier.config.js
export default {
  plugins: ['prettier-plugin-tailwindcss'],
  // v4.0 支持 CSS 配置
  tailwindConfig: './src/styles/main.css',
  // 或传统配置文件
  // tailwindConfig: './tailwind.config.js',
  tailwindFunctions: ['clsx', 'cva', 'cn'],
  tailwindAttributes: ['class', 'className', 'class:list'],
}
```

### 3. **Vite 开发环境优化**

```typescript
// vite.config.ts
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
  plugins: [
    react(),
    tailwindcss(),
  ],
  css: {
    devSourcemap: true, // 开发环境源码映射
  },
  build: {
    cssCodeSplit: true, // CSS 代码分割
    rollupOptions: {
      output: {
        manualChunks: {
          // 将 Tailwind 基础样式单独打包
          'tailwind-base': ['tailwindcss'],
        },
      },
    },
  },
  // v4.0 性能优化
  optimizeDeps: {
    include: ['tailwindcss'],
  },
})
```

### 4. **TypeScript 支持**

```typescript
// types/tailwind.d.ts
import { Config } from 'tailwindcss'

declare module 'tailwindcss' {
  interface Config {
    // v4.0 类型支持
    theme?: {
      extend?: {
        // 自定义主题变量类型
        colors?: Record<string, string | Record<string, string>>
        spacing?: Record<string, string>
        // ... 其他配置
      }
    }
  }
}

// 组件 props 类型
export interface ButtonProps {
  variant?: 'primary' | 'secondary' | 'outline' | 'ghost' | 'destructive'
  size?: 'sm' | 'md' | 'lg' | 'icon'
  className?: string
  children: React.ReactNode
}
```

### 5. **性能监控和分析**

```javascript
// scripts/analyze-css.js
import fs from 'fs'
import { PurgeCSS } from 'purgecss'

// v4.0 CSS 分析工具
async function analyzeTailwindCSS() {
  // 读取构建后的 CSS
  const css = fs.readFileSync('./dist/assets/main.css', 'utf8')
  
  // 分析未使用的样式
  const purgeCSSResult = await new PurgeCSS().purge({
    content: ['./dist/**/*.html', './dist/**/*.js'],
    css: [{ raw: css }],
    safelist: [
      // 保护动态生成的类名
      /data-\[.*\]/,
      /\[.*\]/,
    ],
  })
  
  const originalSize = Buffer.byteLength(css, 'utf8')
  const purgedSize = Buffer.byteLength(purgeCSSResult[0].css, 'utf8')
  const savings = ((originalSize - purgedSize) / originalSize * 100).toFixed(2)
  
  console.log(`Original CSS: ${(originalSize / 1024).toFixed(2)} KB`)
  console.log(`Purged CSS: ${(purgedSize / 1024).toFixed(2)} KB`)
  console.log(`Savings: ${savings}%`)
}

analyzeTailwindCSS()
```

### 6. **自动化工作流**

```yaml
# .github/workflows/css-optimization.yml
name: CSS Optimization

on: [push, pull_request]

jobs:
  optimize:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: '20'
          cache: 'npm'
          
      - name: Install dependencies
        run: npm install
        
      - name: Build CSS
        run: npm run build:css
        
      - name: Analyze CSS size
        run: |
          echo "CSS analysis:"
          ls -la dist/assets/*.css
          node scripts/analyze-css.js
          
      - name: Check for unused styles
        run: npm run lint:css
```

```json
{
  "scripts": {
    "dev": "vite",
    "build": "vite build",
    "build:css": "tailwindcss -i ./src/styles/main.css -o ./dist/assets/main.css --minify",
    "lint:css": "stylelint './src/**/*.css' --fix",
    "analyze:css": "node scripts/analyze-css.js",
    "format": "prettier --write './src/**/*.{js,ts,jsx,tsx,css}'",
    "format:check": "prettier --check './src/**/*.{js,ts,jsx,tsx,css}'"
  }
}
```

## 📱 v4.0 实际项目案例

### 1. **现代电商产品卡片**

```jsx
// ✅ 使用 v4.0 新特性的产品卡片
function ProductCard({ product, onAddToCart }) {
  return (
    <article className="@container group relative">
      <div className="
        bg-[var(--card)] 
        text-[var(--card-foreground)]
        rounded-2xl 
        border 
        border-[var(--border)]
        overflow-hidden 
        transition-all 
        duration-300 
        hover:shadow-xl 
        hover:border-[var(--border)]/80
        transform-gpu
        hover:scale-[1.02]
      ">
        {/* 产品图片区域 */}
        <div className="relative aspect-square overflow-hidden bg-[var(--muted)]">
          <img
            src={product.image}
            alt={product.name}
            className="
              h-full w-full 
              object-cover 
              transition-transform 
              duration-500 
              group-hover:scale-110
            "
          />
          
          {/* 标签 */}
          {product.sale && (
            <div className="
              absolute top-3 left-3 
              bg-[var(--destructive)] 
              text-[var(--destructive-foreground)]
              text-xs font-medium 
              px-3 py-1 
              rounded-full
              animate-bounce
            ">
              Sale
            </div>
          )}
          
          {/* 3D 悬浮按钮 */}
          <button className="
            absolute top-3 right-3 
            p-2 
            bg-[var(--background)]/80 
            backdrop-blur-sm 
            rounded-full 
            opacity-0 
            group-hover:opacity-100 
            transition-all 
            duration-300 
            hover:bg-[var(--background)]
            transform-3d
            hover:translate-z-2
            hover:rotate-y-12
          ">
            <HeartIcon className="w-5 h-5 text-[var(--muted-foreground)]" />
          </button>
          
          {/* 颜色选择器 */}
          <div className="
            absolute bottom-3 left-3 
            flex gap-2
            opacity-0 
            group-hover:opacity-100 
            transition-all 
            duration-300
            starting:translate-y-2
            starting:opacity-0
          ">
            {product.colors?.map((color, index) => (
              <button
                key={color}
                style={{ backgroundColor: color }}
                className={`
                  w-6 h-6 
                  rounded-full 
                  border-2 
                  border-white 
                  shadow-sm
                  transition-all 
                  duration-200
                  hover:scale-110
                  delay-[${index * 50}ms]
                `}
              />
            ))}
          </div>
        </div>
        
        {/* 产品信息 */}
        <div className="p-6 space-y-4">
          {/* 标题和描述 */}
          <div className="space-y-2">
            <h3 className="
              font-semibold 
              text-lg
              text-[var(--foreground)]
              line-clamp-2 
              group-hover:text-[var(--primary)]
              transition-colors
              duration-200
            ">
              {product.name}
            </h3>
            <p className="
              text-sm 
              text-[var(--muted-foreground)]
              line-clamp-2
            ">
              {product.description}
            </p>
          </div>
          
          {/* 评分 - 容器查询响应式 */}
          <div className="flex items-center gap-2 @sm:gap-3">
            <div className="flex items-center">
              {[...Array(5)].map((_, i) => (
                <StarIcon
                  key={i}
                  className={`
                    w-4 h-4 
                    @sm:w-5 
                    @sm:h-5
                    transition-colors
                    duration-200
                    ${i < product.rating
                      ? 'text-yellow-400 fill-current'
                      : 'text-[var(--muted)]'
                    }
                  `}
                />
              ))}
            </div>
            <span className="text-sm text-[var(--muted-foreground)]">
              ({product.reviewCount})
            </span>
          </div>
          
          {/* 价格和购买按钮 */}
          <div className="flex items-center justify-between pt-2">
            <div className="flex items-center gap-2">
              <span className="
                text-xl 
                @sm:text-2xl 
                font-bold 
                text-[var(--foreground)]
                bg-linear-45 
                from-[var(--primary)] 
                to-[var(--accent)]
                bg-clip-text
                text-transparent
              ">
                ${product.price}
              </span>
              {product.originalPrice && (
                <span className="
                  text-sm 
                  text-[var(--muted-foreground)]
                  line-through
                ">
                  ${product.originalPrice}
                </span>
              )}
            </div>
            
            <button
              onClick={() => onAddToCart(product)}
              disabled={!product.inStock}
              className="
                px-4 py-2 
                @sm:px-6 
                @sm:py-3
                bg-[var(--primary)] 
                text-[var(--primary-foreground)]
                text-sm 
                @sm:text-base
                font-medium 
                rounded-lg 
                transition-all 
                duration-200
                hover:bg-[var(--primary)]/90
                focus:outline-none 
                focus:ring-2 
                focus:ring-[var(--ring)]
                focus:ring-offset-2
                disabled:opacity-50 
                disabled:cursor-not-allowed
                transform-gpu
                hover:scale-105
                active:scale-95
                not-disabled:hover:shadow-lg
              "
            >
              {product.inStock ? 'Add to Cart' : 'Out of Stock'}
            </button>
          </div>
        </div>
      </div>
    </article>
  );
}
```

### 2. **v4.0 响应式导航栏**

```jsx
// ✅ 使用 v4.0 特性的现代导航栏
function Navbar() {
  const [isOpen, setIsOpen] = useState(false);
  
  return (
    <nav className="
      sticky top-0 z-50 
      bg-[var(--background)]/80 
      backdrop-blur-xl 
      border-b 
      border-[var(--border)]
      supports-backdrop-blur:bg-[var(--background)]/60
      not-supports-backdrop-blur:bg-[var(--background)]
    ">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          {/* Logo */}
          <div className="flex items-center space-x-3">
            <div className="
              w-8 h-8 
              bg-linear-45/oklch 
              from-[var(--primary)] 
              to-[var(--accent)]
              rounded-lg
              flex items-center justify-center
              transform-3d
              hover:rotate-y-180
              transition-transform
              duration-500
            ">
              <span className="text-white font-bold text-sm">L</span>
            </div>
            <span className="
              text-xl font-bold 
              text-[var(--foreground)]
              hidden sm:block
              bg-linear-to-r/oklch
              from-[var(--foreground)]
              to-[var(--muted-foreground)]
              bg-clip-text
            ">
              Your Brand
            </span>
          </div>
          
          {/* Desktop Menu */}
          <div className="hidden md:flex items-center space-x-8">
            <NavLink href="/" active>Home</NavLink>
            <NavLink href="/products">Products</NavLink>
            <NavLink href="/about">About</NavLink>
            <NavLink href="/contact">Contact</NavLink>
          </div>
          
          {/* Desktop Actions */}
          <div className="hidden md:flex items-center space-x-4">
            <SearchButton />
            <CartButton />
            <ThemeToggle />
            <UserMenu />
          </div>
          
          {/* Mobile menu button */}
          <button
            onClick={() => setIsOpen(!isOpen)}
            className="
              md:hidden
              p-2 
              rounded-lg 
              text-[var(--muted-foreground)]
              hover:text-[var(--foreground)]
              hover:bg-[var(--muted)]
              focus:outline-none 
              focus:ring-2 
              focus:ring-[var(--ring)]
              transition-all
              duration-200
            "
          >
            <div className="relative w-6 h-6">
              {/* 汉堡菜单动画 */}
              <span className={`
                absolute block w-6 h-0.5 
                bg-current 
                transition-all 
                duration-300 
                ease-bounce
                ${isOpen ? 'top-3 rotate-45' : 'top-1'}
              `} />
              <span className={`
                absolute block w-6 h-0.5 
                bg-current 
                top-3 
                transition-all 
                duration-300
                ${isOpen ? 'opacity-0' : 'opacity-100'}
              `} />
              <span className={`
                absolute block w-6 h-0.5 
                bg-current 
                transition-all 
                duration-300 
                ease-bounce
                ${isOpen ? 'top-3 -rotate-45' : 'top-5'}
              `} />
            </div>
          </button>
        </div>
      </div>
      
      {/* Mobile Menu */}
      <div className={`
        md:hidden 
        transition-all 
        duration-300 
        ease-out
        ${isOpen 
          ? 'max-h-screen opacity-100' 
          : 'max-h-0 opacity-0'
        } 
        overflow-hidden 
        bg-[var(--background)]
        border-t 
        border-[var(--border)]
        backdrop-blur-xl
      `}>
        <div className="px-4 pt-4 pb-6 space-y-3">
          {/* 移动端导航链接 */}
          <MobileNavLink href="/" delay="0ms">Home</MobileNavLink>
          <MobileNavLink href="/products" delay="100ms">Products</MobileNavLink>
          <MobileNavLink href="/about" delay="200ms">About</MobileNavLink>
          <MobileNavLink href="/contact" delay="300ms">Contact</MobileNavLink>
          
          {/* 移动端操作按钮 */}
          <div className="
            border-t 
            border-[var(--border)]
            pt-4 
            mt-4
            flex 
            items-center 
            justify-between
          ">
            <SearchButton mobile />
            <CartButton mobile />
            <ThemeToggle />
            <UserMenu mobile />
          </div>
        </div>
      </div>
    </nav>
  );
}

// 导航链接组件
function NavLink({ href, active, children }) {
  return (
    <a
      href={href}
      className={`
        text-sm font-medium 
        transition-all 
        duration-200
        hover:text-[var(--foreground)]
        relative
        group
        ${active 
          ? 'text-[var(--primary)]' 
          : 'text-[var(--muted-foreground)]'
        }
      `}
    >
      {children}
      
      {/* 下划线动画 */}
      <span className={`
        absolute bottom-0 left-0 
        w-full h-0.5 
        bg-[var(--primary)]
        transition-all 
        duration-300
        origin-left
        ${active 
          ? 'scale-x-100' 
          : 'scale-x-0 group-hover:scale-x-100'
        }
      `} />
    </a>
  );
}

// 移动端导航链接
function MobileNavLink({ href, children, delay = "0ms" }) {
  return (
    <a
      href={href}
      style={{ '--delay': delay }}
      className="
        block 
        px-3 py-2 
        text-base font-medium 
        text-[var(--muted-foreground)]
        hover:text-[var(--foreground)]
        hover:bg-[var(--muted)]
        rounded-lg
        transition-all 
        duration-200
        delay-[var(--delay)]
        starting:opacity-0
        starting:translate-x-4
      "
    >
      {children}
    </a>
  );
}
```

## ⚡ v4.0 常见问题和解决方案

### 1. **从 v3 迁移到 v4**

```bash
# 使用官方迁移工具
npx @tailwindcss/upgrade

# 手动迁移步骤：
# 1. 更新依赖
npm uninstall tailwindcss
npm install tailwindcss @tailwindcss/vite

# 2. 更新配置文件
# tailwind.config.js → CSS @theme 语法

# 3. 更新 CSS 导入
# @tailwind base; → @import "tailwindcss";
```

```css
/* v3.x 配置迁移 */
/* 旧的 CSS 文件 */
@tailwind base;
@tailwind components;
@tailwind utilities;

/* v4.0 新的 CSS 文件 */
@import "tailwindcss";

@theme {
  /* 将 tailwind.config.js 中的配置迁移到这里 */
  --color-primary-500: oklch(0.65 0.2 260);
  --font-sans: "Inter", system-ui, sans-serif;
  /* ... */
}
```

### 2. **处理类名过长问题**

```jsx
// ❌ v4.0 中仍然要避免
<div className="@container group relative bg-[var(--card)] text-[var(--card-foreground)] rounded-2xl border border-[var(--border)] overflow-hidden transition-all duration-300 hover:shadow-xl hover:border-[var(--border)]/80 transform-gpu hover:scale-[1.02]">

// ✅ v4.0 推荐做法 - 使用 @layer components
@layer components {
  .product-card {
    @apply @container group relative;
    @apply bg-[var(--card)] text-[var(--card-foreground)];
    @apply rounded-2xl border border-[var(--border)];
    @apply overflow-hidden transition-all duration-300;
    @apply hover:shadow-xl hover:border-[var(--border)]/80;
    @apply transform-gpu hover:scale-[1.02];
  }
}

// 在 JSX 中使用
<div className="product-card">
  {/* 内容 */}
</div>

// ✅ 或使用组件抽象 + CVA
import { cva } from 'class-variance-authority';

const cardVariants = cva(
  // 基础样式
  "@container group relative transition-all duration-300 transform-gpu",
  {
    variants: {
      variant: {
        default: "bg-[var(--card)] text-[var(--card-foreground)] border border-[var(--border)]",
        elevated: "bg-[var(--card)] text-[var(--card-foreground)] shadow-lg border-0",
        outlined: "bg-transparent border-2 border-[var(--border)]",
      },
      size: {
        sm: "p-4 rounded-lg",
        md: "p-6 rounded-xl", 
        lg: "p-8 rounded-2xl",
      },
      interactive: {
        true: "hover:shadow-xl hover:scale-[1.02] cursor-pointer",
        false: "",
      }
    },
    defaultVariants: {
      variant: "default",
      size: "md",
      interactive: true,
    }
  }
);

function Card({ variant, size, interactive, className, children, ...props }) {
  return (
    <div 
      className={cardVariants({ variant, size, interactive, className })}
      {...props}
    >
      {children}
    </div>
  );
}
```

### 3. **状态管理和条件类名（v4.0 增强）**

```jsx
// ✅ v4.0 改进的条件类名处理
import { clsx } from 'clsx';

function Alert({ type, children, onClose, icon }) {
  return (
    <div
      className={clsx(
        // 基础样式
        'p-4 rounded-lg border flex items-start gap-3',
        'transition-all duration-200 transform-gpu',
        
        // v4.0 语义化色彩变体
        {
          'bg-[var(--destructive)]/10 border-[var(--destructive)]/20 text-[var(--destructive)]': type === 'error',
          'bg-yellow-50 border-yellow-200 text-yellow-800 dark:bg-yellow-900/20 dark:border-yellow-800/30 dark:text-yellow-400': type === 'warning',
          'bg-green-50 border-green-200 text-green-800 dark:bg-green-900/20 dark:border-green-800/30 dark:text-green-400': type === 'success',
          'bg-blue-50 border-blue-200 text-blue-800 dark:bg-blue-900/20 dark:border-blue-800/30 dark:text-blue-400': type === 'info',
        },
        
        // v4.0 动画支持
        {
          'starting:opacity-0 starting:scale-95 starting:translate-y-2': true,
        }
      )}
    >
      {/* 图标 */}
      {icon && (
        <div className="flex-shrink-0">
          <AlertIcon type={type} />
        </div>
      )}
      
      {/* 内容 */}
      <div className="flex-1 min-w-0">
        {children}
      </div>
      
      {/* 关闭按钮 */}
      {onClose && (
        <button
          onClick={onClose}
          className="
            flex-shrink-0
            text-current/70 
            hover:text-current 
            transition-colors
            duration-200
            hover:bg-current/10
            rounded-md
            p-1
            -m-1
          "
        >
          <XIcon className="w-4 h-4" />
        </button>
      )}
    </div>
  );
}

// ✅ v4.0 数据属性支持
function InteractiveCard({ isActive, isSelected, children }) {
  return (
    <div
      data-active={isActive}
      data-selected={isSelected}
      className="
        p-4 
        border 
        rounded-lg
        transition-all 
        duration-200
        
        /* v4.0 原生数据属性支持 */
        data-active:border-[var(--primary)]
        data-active:bg-[var(--primary)]/5
        data-selected:ring-2
        data-selected:ring-[var(--ring)]
        
        /* 组合状态 */
        data-active:data-selected:bg-[var(--primary)]/10
        
        /* 非状态样式 */
        not-data-active:opacity-75
        not-data-selected:hover:border-[var(--border)]/80
      "
    >
      {children}
    </div>
  );
}
```

### 4. **性能优化技巧**

```css
/* v4.0 性能优化 */
@theme {
  /* 使用 CSS 变量减少重复 */
  --transition-fast: 150ms ease-out;
  --transition-normal: 200ms ease-out;
  --transition-slow: 300ms ease-out;
  
  /* 预定义常用组合 */
  --focus-ring: 0 0 0 2px var(--ring);
  --shadow-focus: var(--shadow-md), var(--focus-ring);
}

@layer utilities {
  /* 性能优化的工具类 */
  .transform-gpu {
    transform: translateZ(0);
    will-change: transform;
  }
  
  .contain-layout {
    contain: layout;
  }
  
  .contain-paint {
    contain: paint;
  }
  
  /* 减少重绘的文本渲染 */
  .text-render-optimized {
    text-rendering: optimizeSpeed;
    font-smooth: never;
    -webkit-font-smoothing: none;
  }
}
```

### 5. **调试和开发技巧**

```javascript
// v4.0 调试工具
// scripts/debug-tailwind.js
import fs from 'fs';
import path from 'path';

// 分析 CSS 变量使用情况
function analyzeCSSVariables(cssContent) {
  const variableRegex = /var\(--[\w-]+\)/g;
  const variables = cssContent.match(variableRegex) || [];
  const usage = {};
  
  variables.forEach(variable => {
    const name = variable.slice(4, -1); // 移除 var( 和 )
    usage[name] = (usage[name] || 0) + 1;
  });
  
  return Object.entries(usage)
    .sort(([,a], [,b]) => b - a)
    .slice(0, 10);
}

// 检查未使用的主题变量
function findUnusedThemeVariables(themeVars, usedVars) {
  const unused = [];
  for (const variable of themeVars) {
    if (!usedVars.includes(variable)) {
      unused.push(variable);
    }
  }
  return unused;
}

// 开发模式类名验证
function validateClassNames(classNames) {
  const issues = [];
  
  classNames.forEach(className => {
    // 检查常见错误
    if (className.includes('grid-cols-') && !className.match(/grid-cols-\d+/)) {
      issues.push(`Invalid grid columns: ${className}`);
    }
    
    // 检查 v4.0 特定语法
    if (className.includes('@') && !className.match(/@(sm|md|lg|xl|2xl|container|min-|max-)/)) {
      issues.push(`Invalid container query: ${className}`);
    }
  });
  
  return issues;
}

console.log('TailwindCSS v4.0 Debug Report');
```
```

## 📋 团队协作规范

### 1. **命名约定**

```javascript
// 组件样式变体命名
const buttonVariants = {
  // 主要变体
  primary: 'bg-blue-500 text-white hover:bg-blue-600',
  secondary: 'bg-gray-200 text-gray-900 hover:bg-gray-300',
  outline: 'border border-gray-300 text-gray-700 hover:bg-gray-50',
  
  // 尺寸变体
  sm: 'px-3 py-1.5 text-sm',
  md: 'px-4 py-2 text-base',
  lg: 'px-6 py-3 text-lg',
  
  // 状态变体
  loading: 'opacity-50 cursor-not-allowed',
  disabled: 'opacity-50 cursor-not-allowed',
};
```

### 2. **代码审查检查点**

```markdown
## TailwindCSS代码审查清单

### 样式相关
- [ ] 是否使用了语义化的颜色变量
- [ ] 响应式设计是否完整
- [ ] 是否遵循移动优先原则
- [ ] 暗色主题支持是否完整

### 性能相关
- [ ] 是否避免了动态类名
- [ ] 是否正确配置了purge选项
- [ ] 是否复用了通用组件

### 可维护性
- [ ] 复杂样式是否提取为组件
- [ ] 是否使用了一致的间距系统
- [ ] 类名顺序是否一致
```

## 🎉 总结

TailwindCSS的最佳实践关键在于：

1. **设计系统优先**：建立一致的设计令牌
2. **组件化思维**：抽象复用样式
3. **响应式设计**：移动优先原则
4. **性能优化**：正确配置purge
5. **团队协作**：统一规范和工具

通过这些实践，你可以构建出既美观又高性能的现代Web应用！

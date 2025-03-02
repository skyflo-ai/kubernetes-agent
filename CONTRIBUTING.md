# Contributing to Skyflo Kubernetes Agent

First off, thank you for your interest in contributing to Skyflo Kubernetes Agent! We value your time and effort to help make this project better for everyone.

## Code of Conduct

Please note that this project is governed by a [Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project you agree to abide by its terms.

## How Can I Contribute?

### Reporting Issues

If you find a bug or have a suggestion for an improvement, please open an issue using the [issue template](.github/ISSUE_TEMPLATE.md). Ensure to provide as much detail as possible to help us understand and resolve it.

### Suggesting Enhancements

We welcome feature requests and enhancements. When suggesting a feature, please describe the problem you're trying to solve and how you think the feature would address it.

### Your First Code Contribution

Unsure where to start? Check out our [good first issues](https://github.com/your_username/skyflo-kubernetes-agent/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) to find tasks that are suitable for newcomers.

## Pull Requests

1. **Fork the repository.**
2. **Create a feature branch:**

    ```bash
    git checkout -b feature/YourFeatureName
    ```

3. **Commit your changes:**

    ```bash
    git commit -m 'Add some feature'
    ```

4. **Push to the branch:**

    ```bash
    git push origin feature/YourFeatureName
    ```

5. **Open a pull request.**

Please ensure your pull request adheres to the following guidelines:

- **Provide a clear description:** Explain the changes you're making and the reasons behind them.
- **Link to relevant issues:** If your pull request addresses an existing issue, please link to it.
- **Ensure code quality:** Follow the [Coding Guidelines](#coding-guidelines) below.

## Coding Guidelines

- **Language:** This project is written in Go. Follow the [Go Code Review Guidelines](https://github.com/golang/go/wiki/CodeReviewComments).
- **Formatting:** Ensure code is properly formatted. Use `gofmt` to format your code.
- **Documentation:** Document your code where necessary. Update any relevant documentation if you add a function or feature.
- **Testing:** Write unit and integration tests for your changes. Ensure all tests pass before submitting your pull request.

## Style Guide

- **Commit Messages:** Write clear, concise commit messages. Follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification.
- **Naming Conventions:** Use descriptive names for variables, functions, and packages.
- **Code Structure:** Organize code logically and adhere to the project's existing structure.

## Testing Your Changes

Before submitting a pull request, make sure your changes pass all tests:

```bash
make test
```

## Continuous Integration

All pull requests are automatically tested using our CI pipeline. Ensure that your pull request does not break any existing tests and passes all CI checks before merging.

## Development Setup

To set up a local development environment, follow these steps:

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/your_username/skyflo-kubernetes-agent.git
    cd skyflo-kubernetes-agent
    ```

2. **Install Dependencies:**

    Ensure you have all prerequisites installed as outlined in the [README.md](README.md).

3. **Build the Project:**

    ```bash
    make build
    ```

4. **Run Tests:**

    ```bash
    make test
    ```

## Communication

For questions, discussions, and support, please join our [community channels](https://github.com/your_username/skyflo-kubernetes-agent#communication).

## License

By contributing, you agree that your contributions will be licensed under the [MIT License](LICENSE). 
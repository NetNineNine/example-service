# Contributing to Example Project

Thank you for considering contributing to our project! To make the process as smooth as possible, please follow these guidelines.

## Process

1. **Create a Branch Linked to an OpenProject Work Item**

   Every change must be linked to an OpenProject work item (e.g., task, bug, feature). Create a new branch with a meaningful name that includes the work item ID.

   ```sh
   git checkout -b workitem-123-meaningful-description
   ```

2. **Make Changes in the Template Folder**

   Make all necessary changes in the `template` folder. This is where the reusable code and files for the project reside.

   ```sh
   cd template
   # Make your changes
   ```

3. **Run the Test Script**

   After making your changes, you need to verify that everything works as expected by running the test script. This script copies the template to a test directory and runs any necessary checks.

   ```sh
   ./copy-test.sh
   ```

   Ensure that the script completes successfully without any errors. It will update the root of this project so you can test the CI.

4. **Commit Your Changes**

   Once you're satisfied with your changes, add and commit them. Use Conventional Commits (see [cheat sheet](https://gist.github.com/qoomon/5dfcdf8eec66a051ecd85625518cfd13)).

   ```sh
   git add -A
   git commit -m "fix: use the correct logger"
   ```

5. **Push Your Branch**

   Push your branch to the origin repository.

   ```sh
   git push origin workitem-123-meaningful-description
   ```

6. **Log a Pull Request (PR)**

   Open a pull request against the main branch. Ensure your PR description includes a clear explanation of your changes and references the OpenProject work item.

   1. Go to the repository on GitHub.
   2. Click on the "Pull requests" tab.
   3. Click the "New pull request" button.
   4. Select your branch from the dropdown menu.
   5. Provide a descriptive title and detailed description for your PR.
   6. Click the "Create pull request" button.
   7. Fill in the template details.
   8. Go through the checklist and ensure everything is done before others look at it.

Thank you for your contributions! Please reach out if you have any questions or need further assistance.
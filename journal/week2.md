# Week 2 - Language Learning App With Python


This entry provides an overview of the steps taken to set up and begin development of the Language Learning Assistant project.

**Summary:** This section outlines the initial steps taken to set up the project, including  creating and activating a virtual environment, and installing initial dependencies.

## Project Setup


1.  **Created a Python virtual environment** named `venv` to manage project dependencies in isolation:
    ```bash
    python -m venv venv
    ```

2.  **Activated the virtual environment** :
    ```bash
    source venv/bin/activate
    ```

## Dependency Management

1.  **Installed initial dependencies** listed in `requirements.txt` using pip:
    ```bash
    pip install -r requirements.txt
    ```
    This included libraries such as `chromadb`, `streamlit`, and `boto3`.

2.  Later, **`youtube_transcript_api` was added to `requirements.txt`** and installed:
    ```bash
    pip install youtube_transcript_api
    ```

**Summary:** This section details the steps taken to manage project dependencies, including the initial installation of required libraries and the subsequent addition of the `youtube_transcript_api`.

## Backend Implementation

1.  **Navigated to the `backend` directory** to work on the server-side logic:
    ```bash
    cd backend
    ```

2.  **Created a `requirements.txt` file** within the `backend` directory to list backend-specific dependencies, starting with `chromadb`.

3.  **Attempted to run `rag.py`**, which encountered an **`sqlite3` version error**. ChromaDB required version 3.35.0 or higher, but the system had 3.31.0. Efforts to upgrade `sqlite3` via `apt` were made to resolve the bottleneck.

4.  **Investigated alternative vector store solutions** such as PG Vector and ChromaDB's self-hosted option. The open-source version of ChromaDB was chosen for its ease of installation via pip.

5.  **Integrated chat functionality** via `chat.py`.

6.  **Tested the basic chat functionality** by running `chat.py` independently:
    ```bash
    python chat.py
    ```
    This allowed interaction with Amazon Bedrock models (e.g., Nova micro).

7.  **Integrated YouTube transcript downloading functionality**.

8.  **Modified `backend/transcripts/get_transcripts.py`** to include a `save` method for storing downloaded transcripts in a `transcripts` folder.

**Summary:** This section describes the development of the backend, including resolving dependency issues, integrating chat and transcript downloading functionalities, and modifying code for transcript storage.

## Frontend Implementation

1.  **Navigated to the `frontend` directory** to work on the user interface :
    ```bash
    cd frontend
    ```

2.  **Attempted to run the Streamlit front end** using:
    ```bash
    streamlit run main.py
    ```
    This initially resulted in an **`ImportError: No module named 'backend'`**.

3.  **Troubleshooting the import error** involved trying different run commands and import statement modifications.

4.  The import issue was temporarily addressed by adding the following import statement in `frontend/main.py`:
    ```python
    import sys
    import os
    sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
    ```

5.  **Tested the Streamlit application** on `http://localhost:8501`. The application features different stages: Chat, Raw Text (for downloading transcripts), Structure Data, RAG Implementation, and Interactive Learning.

6.  The **"Raw Transcript" stage was successfully tested** by pasting a YouTube video URL and downloading its transcript using the integrated `youtube_transcript_api`. The downloaded transcript was displayed in the application.

**Summary:** This section details the development of the frontend, including troubleshooting import errors, testing the Streamlit application, and successfully implementing the "Raw Transcript" stage for downloading YouTube transcripts.

## Version Control

1.  **Utilized Git for version control**, including commands like `git clone`, `git add`, `git commit`, and `git pull` to manage code changes and integrate updates from remote repositories.


**Summary:** This section highlights the use of Git for version control and collaboration, including merging pull requests.

## Functionality Testing

1.  **Basic chat functionality with Amazon Bedrock models was tested** through the independent `chat.py` script.

2.  **YouTube transcript downloading was successfully implemented and tested** through the Streamlit front end using the "Raw Transcript" stage.

3.  **Initial analysis of a downloaded transcript's structure** was performed by pasting it into the chat interface and querying the language model.

**Summary:** This section summarizes the testing of core functionalities, including chat and transcript downloading, and initial analysis of transcript structures.

## Further Steps

The next steps involve further developing the backend to structure the downloaded transcripts, integrate them into the ChromaDB vector store, implement the Retrieval-Augmented Generation (RAG) pipeline, and build out the interactive learning features in the Streamlit front end.

**Summary:** This section outlines the planned next steps for the project, focusing on enhancing backend functionality and developing interactive learning features.



# Technical Challenges Faced While Setting Up Environment and Installing Packages

During the process of setting up my Python environment and installing packages, I encountered several challenges. Here’s a step-by-step record of what I did to troubleshoot and resolve the issues.

**Summary:** This section introduces the technical challenges faced during environment setup and package installation.

## 1. Setting Up the Python Virtual Environment (venv)

1. First, I created a Python virtual environment:
    ```bash
    python3.12 -m venv venv
    ```
    After that, I activated the virtual environment:
    ```bash
    source venv/bin/activate  # For Linux/macOS
    ```

**Summary:** This section details the initial setup of the Python virtual environment.

## 2. Installing Packages from requirements.txt

I ran the following command to install the packages from `requirements.txt`:
    ```bash
    pip install -r requirements.txt
    ```
However, the installation process was hanging, so I started troubleshooting.

**Summary:** This section describes the initial attempt to install packages and the resulting issue.

## 3. Step-by-Step Troubleshooting

### 3.1 Testing Simple Package Installation

To verify if `pip` was working, I tried installing a simple package (`requests`):
    ```bash
    pip install requests
    ```
This command also hung, indicating that the issue was likely with the network or `pip`.

### 3.2 Running pip with Verbose Output

To get more information about what was happening during the installation, I ran the command with `-v` (verbose mode):
    ```bash
    pip install -r requirements.txt -v --no-cache-dir --disable-pip-version-check
    ```
This output showed me that the process was hanging while downloading metadata for a package (`chromadb`). This helped me identify where the issue was.

### 3.3 Testing Network Connection

To check if there was a network issue, I tried pinging `pypi.org`:
    ```bash
    ping -4 pypi.org
    ```
This worked fine, so the issue wasn’t with my connection to PyPI.

**Summary:** This section outlines the initial troubleshooting steps, including testing simple package installations, using verbose output, and checking network connectivity.

## 4. Changing DNS to Resolve Network Issues

Since the installation was still hanging, I suspected that there might be a DNS issue. I decided to change my DNS settings.

I opened the `/etc/resolv.conf` file:
    ```bash
    sudo nano /etc/resolv.conf
    ```
Then, I replaced the existing DNS servers with Google's public DNS servers:
    ```nginx
    nameserver 8.8.8.8
    nameserver 8.8.4.4
    ```
After saving the changes, I ran the `pip install` command again, and it worked without any issues.

**Summary:** This section describes the resolution of network issues by changing DNS settings.


